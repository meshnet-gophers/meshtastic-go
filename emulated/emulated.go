package emulated

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/meshnet-gophers/meshtastic-go"
	pb "github.com/meshnet-gophers/meshtastic-go/meshtastic"
	"github.com/meshnet-gophers/meshtastic-go/mqtt"
	"github.com/meshnet-gophers/meshtastic-go/radio"
	"github.com/meshnet-gophers/meshtastic-go/transport"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"sync"
	"time"
)

const (
	// MinAppVersion is the minimum app version supported by the emulated radio.
	MinAppVersion = 30200
)

// Config is the configuration for the emulated Radio.
type Config struct {
	// Dependencies
	MQTTClient *mqtt.Client

	// Node configuration
	// NodeID is the ID of the node.
	NodeID meshtastic.NodeID
	// LongName is the long name of the node.
	LongName string
	// ShortName is the short name of the node.
	ShortName string
	// Channels is the set of channels the radio will listen and transmit on.
	// The first channel in the set is considered the primary channel and is used for broadcasting NodeInfo and Position
	Channels *pb.ChannelSet
	// BroadcastNodeInfoInterval is the interval at which the radio will broadcast a NodeInfo on the Primary channel.
	// The zero value disables broadcasting NodeInfo.
	BroadcastNodeInfoInterval time.Duration

	// BroadcastPositionInterval is the interval at which the radio will broadcast Position on the Primary channel.
	// The zero value disables broadcasting NodeInfo.
	BroadcastPositionInterval time.Duration
	// PositionLatitudeI is the latitude of the position which will be regularly broadcasted.
	// This is in degrees multiplied by 1e7.
	PositionLatitudeI int32
	// PositionLongitudeI is the longitude of the position which will be regularly broadcasted.
	// This is in degrees multiplied by 1e7.
	PositionLongitudeI int32
	// PositionAltitude is the altitude of the position which will be regularly broadcasted.
	// This is in meters above MSL.
	PositionAltitude int32

	// TCPListenAddr is the address the emulated radio will listen on for TCP connections and offer the Client API over.
	TCPListenAddr string
}

func (c *Config) validate() error {
	if c.MQTTClient == nil {
		return fmt.Errorf("MQTTClient is required")
	}
	if c.NodeID == 0 {
		return fmt.Errorf("NodeID is required")
	}
	if c.LongName == "" {
		c.LongName = c.NodeID.DefaultLongName()
	}
	if c.ShortName == "" {
		c.ShortName = c.NodeID.DefaultShortName()
	}
	if c.Channels == nil {
		//lint:ignore ST1005 we're referencing an actual field here.
		return fmt.Errorf("Channels is required")
	}
	if len(c.Channels.Settings) == 0 {
		return fmt.Errorf("Channels.Settings should be non-empty")
	}
	return nil
}

// Radio emulates a meshtastic Node, communicating with a meshtastic network via MQTT.
type Radio struct {
	cfg    Config
	mqtt   *mqtt.Client
	logger *log.Logger

	// TODO: rwmutex?? seperate mutexes??
	mu                   sync.Mutex
	fromRadioSubscribers map[chan<- *pb.FromRadio]struct{}
	nodeDB               map[uint32]*pb.NodeInfo
	// packetID is incremented and included in each packet sent from the radio.
	// TODO: Eventually, we should offer an easy way of persisting this so that we can resume from where we left off.
	packetID uint32
}

// NewRadio creates a new emulated radio.
func NewRadio(cfg Config) (*Radio, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}
	return &Radio{
		cfg:                  cfg,
		logger:               log.With("radio", cfg.NodeID.String()),
		fromRadioSubscribers: map[chan<- *pb.FromRadio]struct{}{},
		mqtt:                 cfg.MQTTClient,
		nodeDB:               map[uint32]*pb.NodeInfo{},
	}, nil
}

// Run starts the radio. It blocks until the context is cancelled.
func (r *Radio) Run(ctx context.Context) error {
	if err := r.mqtt.Connect(); err != nil {
		return fmt.Errorf("connecting to mqtt: %w", err)
	}
	// TODO: Disconnect??

	// Subscribe to all configured channels
	for _, ch := range r.cfg.Channels.Settings {
		r.logger.Debug("subscribing to mqtt for channel", "channel", ch.Name)
		r.mqtt.Handle(ch.Name, r.handleMQTTMessage)
	}

	// TODO: Rethink concurrency. Do we want a goroutine servicing ToRadio and one servicing FromRadio?

	eg, egCtx := errgroup.WithContext(ctx)
	// Spin up goroutine to send NodeInfo every interval
	if r.cfg.BroadcastNodeInfoInterval > 0 {
		eg.Go(func() error {
			ticker := time.NewTicker(r.cfg.BroadcastNodeInfoInterval)
			defer ticker.Stop()
			for {
				if err := r.broadcastNodeInfo(egCtx); err != nil {
					r.logger.Error("failed to broadcast node info", "err", err)
				}
				select {
				case <-egCtx.Done():
					return nil
				case <-ticker.C:
				}
			}
		})
	}
	// Spin up goroutine to send Position every interval
	if r.cfg.BroadcastPositionInterval > 0 {
		eg.Go(func() error {
			ticker := time.NewTicker(r.cfg.BroadcastPositionInterval)
			defer ticker.Stop()
			for {
				if err := r.broadcastPosition(egCtx); err != nil {
					r.logger.Error("failed to broadcast position", "err", err)
				}
				select {
				case <-egCtx.Done():
					return nil
				case <-ticker.C:
				}
			}
		})
	}
	if r.cfg.TCPListenAddr != "" {
		eg.Go(func() error {
			return r.listenTCP(egCtx)
		})
	}

	return eg.Wait()
}

func (r *Radio) handleMQTTMessage(msg mqtt.Message) {
	// TODO: Determine how "github.com/eclipse/paho.mqtt.golang" handles concurrency. Do we need to dispatch here to
	// a goroutine which handles incoming messages to unblock this one?
	if err := r.tryHandleMQTTMessage(msg); err != nil {
		r.logger.Error("failed to handle incoming mqtt message", "err", err)
	}
}

func (r *Radio) updateNodeDB(nodeID uint32, updateFunc func(*pb.NodeInfo)) {
	r.mu.Lock()
	defer r.mu.Unlock()
	nodeInfo, ok := r.nodeDB[nodeID]
	if !ok {
		nodeInfo = &pb.NodeInfo{
			Num: nodeID,
		}
	}
	updateFunc(nodeInfo)
	nodeInfo.LastHeard = uint32(time.Now().Unix())
	r.nodeDB[nodeID] = nodeInfo
}

func (r *Radio) getNodeDB() []*pb.NodeInfo {
	r.mu.Lock()
	defer r.mu.Unlock()
	nodes := make([]*pb.NodeInfo, 0, len(r.nodeDB))
	for _, node := range r.nodeDB {
		clonedNode := proto.Clone(node).(*pb.NodeInfo)
		nodes = append(nodes, clonedNode)
	}
	return nodes
}

func (r *Radio) tryHandleMQTTMessage(msg mqtt.Message) error {
	serviceEnvelope := &pb.ServiceEnvelope{}
	if err := proto.Unmarshal(msg.Payload, serviceEnvelope); err != nil {
		return fmt.Errorf("unmarshalling: %w", err)
	}
	meshPacket := serviceEnvelope.Packet

	// TODO: Attempt decryption first before dispatching to subscribers
	// TODO: This means we move this further below.
	if err := r.dispatchMessageToFromRadio(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_Packet{
			Packet: meshPacket,
		},
	}); err != nil {
		r.logger.Error("failed to dispatch message to FromRadio subscribers", "err", err)
	}

	// From now on, we only care about messages on the primary channel
	primaryName := r.cfg.Channels.Settings[0].Name
	primaryPSK := r.cfg.Channels.Settings[0].Psk
	if serviceEnvelope.ChannelId != primaryName {
		return nil
	}

	r.logger.Debug("received service envelope for primary channel", "serviceEnvelope", serviceEnvelope)
	// Check if we should try and decrypt the message
	var data *pb.Data
	switch payload := meshPacket.PayloadVariant.(type) {
	case *pb.MeshPacket_Decoded:
		data = payload.Decoded
	case *pb.MeshPacket_Encrypted:
		// TODO: Check if we have the key for this channel
		plaintext, err := radio.XOR(
			payload.Encrypted,
			primaryPSK,
			meshPacket.Id,
			meshPacket.From,
		)
		if err != nil {
			return fmt.Errorf("decrypting: %w", err)
		}
		data = &pb.Data{}
		if err := proto.Unmarshal(plaintext, data); err != nil {
			return fmt.Errorf("unmarshalling decrypted data: %w", err)
		}
	default:
		return fmt.Errorf("unknown payload variant %T", payload)
	}
	r.logger.Debug("received data for primary channel", "data", data)

	// For messages on the primary channel, we want to handle these and potentially update the nodeDB.
	switch data.Portnum {
	case pb.PortNum_NODEINFO_APP:
		user := &pb.User{}
		if err := proto.Unmarshal(data.Payload, user); err != nil {
			return fmt.Errorf("unmarshalling user: %w", err)
		}
		r.logger.Info("received NodeInfo", "user", user)
		r.updateNodeDB(meshPacket.From, func(nodeInfo *pb.NodeInfo) {
			nodeInfo.User = user
		})
	case pb.PortNum_TEXT_MESSAGE_APP:
		r.logger.Info("received TextMessage", "message", string(data.Payload))
	case pb.PortNum_ROUTING_APP:
		routingPayload := &pb.Routing{}
		if err := proto.Unmarshal(data.Payload, routingPayload); err != nil {
			return fmt.Errorf("unmarshalling routingPayload: %w", err)
		}
		r.logger.Info("received Routing", "routing", routingPayload)
	case pb.PortNum_POSITION_APP:
		positionPayload := &pb.Position{}
		if err := proto.Unmarshal(data.Payload, positionPayload); err != nil {
			return fmt.Errorf("unmarshalling positionPayload: %w", err)
		}
		r.logger.Info("received Position", "position", positionPayload)
		r.updateNodeDB(meshPacket.From, func(nodeInfo *pb.NodeInfo) {
			nodeInfo.Position = positionPayload
		})
	case pb.PortNum_TELEMETRY_APP:
		telemetryPayload := &pb.Telemetry{}
		if err := proto.Unmarshal(data.Payload, telemetryPayload); err != nil {
			return fmt.Errorf("unmarshalling telemetryPayload: %w", err)
		}
		deviceMetrics := telemetryPayload.GetDeviceMetrics()
		if deviceMetrics == nil {
			break
		}
		r.logger.Info("received Telemetry deviceMetrics", "telemetry", telemetryPayload)
		r.updateNodeDB(meshPacket.From, func(nodeInfo *pb.NodeInfo) {
			nodeInfo.DeviceMetrics = deviceMetrics
		})
	default:
		r.logger.Debug("received unhandled app payload", "data", data)
	}

	return nil
}

func (r *Radio) nextPacketID() uint32 {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.packetID++
	return r.packetID
}

func (r *Radio) sendPacket(ctx context.Context, packet *pb.MeshPacket) error {
	// TODO: Optimistically attempt to encrypt the packet here if we recognise the channel, encryption is enabled and
	// the payload is not currently encrypted.

	// sendPacket is responsible for setting the packet ID.
	r.packetID = r.nextPacketID()

	se := &pb.ServiceEnvelope{
		// TODO: Fetch channel to use based on packet.Channel rather than hardcoding to primary channel.
		ChannelId: r.cfg.Channels.Settings[0].Name,
		GatewayId: r.cfg.NodeID.String(),
		Packet:    packet,
	}
	bytes, err := proto.Marshal(se)
	if err != nil {
		return fmt.Errorf("marshalling service envelope: %w", err)
	}
	return r.mqtt.Publish(&mqtt.Message{
		Topic:   r.mqtt.GetFullTopicForChannel(r.cfg.Channels.Settings[0].Name) + "/" + r.cfg.NodeID.String(),
		Payload: bytes,
	})
}

func (r *Radio) broadcastNodeInfo(ctx context.Context) error {
	r.logger.Info("broadcasting NodeInfo")
	// TODO: Lots of stuff missing here. However, this is enough for it to show in the UI of another node listening to
	// the MQTT server.
	user := &pb.User{
		Id:        r.cfg.NodeID.String(),
		LongName:  r.cfg.LongName,
		ShortName: r.cfg.ShortName,
		HwModel:   pb.HardwareModel_PRIVATE_HW,
	}
	userBytes, err := proto.Marshal(user)
	if err != nil {
		return fmt.Errorf("marshalling user: %w", err)
	}
	return r.sendPacket(ctx, &pb.MeshPacket{
		From: r.cfg.NodeID.Uint32(),
		To:   meshtastic.BroadcastNodeID.Uint32(),
		PayloadVariant: &pb.MeshPacket_Decoded{
			Decoded: &pb.Data{
				Portnum: pb.PortNum_NODEINFO_APP,
				Payload: userBytes,
			},
		},
	})
}

func (r *Radio) broadcastPosition(ctx context.Context) error {
	r.logger.Info("broadcasting Position")
	position := &pb.Position{
		LatitudeI:  r.cfg.PositionLatitudeI,
		LongitudeI: r.cfg.PositionLongitudeI,
		Altitude:   r.cfg.PositionAltitude,
		Time:       uint32(time.Now().Unix()),
	}
	positionBytes, err := proto.Marshal(position)
	if err != nil {
		return fmt.Errorf("marshalling position: %w", err)
	}
	return r.sendPacket(ctx, &pb.MeshPacket{
		From: r.cfg.NodeID.Uint32(),
		To:   meshtastic.BroadcastNodeID.Uint32(),
		PayloadVariant: &pb.MeshPacket_Decoded{
			Decoded: &pb.Data{
				Portnum: pb.PortNum_POSITION_APP,
				Payload: positionBytes,
			},
		},
	})
}

// dispatchMessageToFromRadio sends a FromRadio message to all current subscribers to
// the FromRadio.
func (r *Radio) dispatchMessageToFromRadio(msg *pb.FromRadio) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for ch := range r.fromRadioSubscribers {
		// TODO: Make this way safer/resilient
		ch <- msg
	}
	return nil
}

func (r *Radio) handleToRadioWantConfigID(conn *transport.StreamConn, req *pb.ToRadio_WantConfigId) error {
	// Send MyInfo
	err := conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_MyInfo{
			MyInfo: &pb.MyNodeInfo{
				MyNodeNum:   r.cfg.NodeID.Uint32(),
				RebootCount: 0,
				// TODO: Track this as a const
				MinAppVersion: MinAppVersion,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}

	// Send Metadata
	err = conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_Metadata{
			Metadata: &pb.DeviceMetadata{
				// TODO: Establish firmwareVersion/deviceStateVersion to fake here
				FirmwareVersion:    "2.2.19-fake",
				DeviceStateVersion: 22,
				CanShutdown:        true,
				HasWifi:            true,
				HasBluetooth:       true,
				// PositionFlags?
				HwModel: pb.HardwareModel_PRIVATE_HW,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}

	// Send all NodeDB entries - plus myself.
	// TODO: Our own node info entry should be in the DB to avoid the special case here.
	err = conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_NodeInfo{
			NodeInfo: &pb.NodeInfo{
				Num: r.cfg.NodeID.Uint32(),
				User: &pb.User{
					Id:        r.cfg.NodeID.String(),
					LongName:  r.cfg.LongName,
					ShortName: r.cfg.ShortName,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}
	for _, nodeInfo := range r.getNodeDB() {
		err = conn.Write(&pb.FromRadio{
			PayloadVariant: &pb.FromRadio_NodeInfo{
				NodeInfo: nodeInfo,
			},
		})
		if err != nil {
			return fmt.Errorf("writing to streamConn: %w", err)
		}
	}

	// TODO: Send all channels
	err = conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_Channel{
			Channel: &pb.Channel{
				Index: 0,
				Settings: &pb.ChannelSettings{
					Psk: nil,
				},
				Role: pb.Channel_PRIMARY,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}

	// Send Config: Device
	err = conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_Config{
			Config: &pb.Config{
				PayloadVariant: &pb.Config_Device{
					Device: &pb.Config_DeviceConfig{
						SerialEnabled:         true,
						NodeInfoBroadcastSecs: uint32(r.cfg.BroadcastNodeInfoInterval.Seconds()),
					},
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}

	// Send ConfigComplete to indicate we're done
	err = conn.Write(&pb.FromRadio{
		PayloadVariant: &pb.FromRadio_ConfigCompleteId{
			ConfigCompleteId: req.WantConfigId,
		},
	})
	if err != nil {
		return fmt.Errorf("writing to streamConn: %w", err)
	}

	return nil
}

func (r *Radio) handleConn(ctx context.Context, underlying io.ReadWriteCloser) error {
	streamConn := transport.NewRadioStreamConn(underlying)
	defer func() {
		if err := streamConn.Close(); err != nil {
			r.logger.Error("failed to close streamConn", "err", err)
		}
	}()

	eg, egCtx := errgroup.WithContext(ctx)
	// Handling messages coming from client
	eg.Go(func() error {
		for {
			select {
			case <-egCtx.Done():
				return nil
			default:
			}
			msg := &pb.ToRadio{}
			if err := streamConn.Read(msg); err != nil {
				return fmt.Errorf("reading from streamConn: %w", err)
			}
			r.logger.Info("received ToRadio from streamConn", "msg", msg)
			switch payload := msg.PayloadVariant.(type) {
			case *pb.ToRadio_Disconnect:
				// The meshtastic python client sends a Disconnect command and with the TCP implementation, it expects
				// the radio to close the connection. So we end the read loop here, and return to close the connection.
				return nil
			case *pb.ToRadio_WantConfigId:
				if err := r.handleToRadioWantConfigID(streamConn, payload); err != nil {
					return fmt.Errorf("handling WantConfigId: %w", err)
				}
			case *pb.ToRadio_Packet:
				if decoded := payload.Packet.GetDecoded(); decoded != nil {
					if decoded.Portnum == pb.PortNum_ADMIN_APP {
						admin := &pb.AdminMessage{}
						if err := proto.Unmarshal(decoded.Payload, admin); err != nil {
							return fmt.Errorf("unmarshalling admin: %w", err)
						}

						switch adminPayload := admin.PayloadVariant.(type) {
						// TODO: Properly handle channel listing, this hack is just so the Python CLI thinks
						// it's connected
						case *pb.AdminMessage_GetChannelRequest:
							r.logger.Info("received GetChannelRequest", "adminPayload", adminPayload, "packet", payload)
							resp := &pb.AdminMessage{
								PayloadVariant: &pb.AdminMessage_GetChannelResponse{
									GetChannelResponse: &pb.Channel{
										Index: 0,
										Settings: &pb.ChannelSettings{
											Psk: nil,
										},
										Role: pb.Channel_DISABLED,
									},
								},
							}
							respBytes, err := proto.Marshal(resp)
							if err != nil {
								return fmt.Errorf("marshalling GetChannelResponse: %w", err)
							}
							// Send GetChannelResponse
							if err := streamConn.Write(&pb.FromRadio{
								PayloadVariant: &pb.FromRadio_Packet{
									Packet: &pb.MeshPacket{
										Id:   r.nextPacketID(),
										From: r.cfg.NodeID.Uint32(),
										To:   r.cfg.NodeID.Uint32(),
										PayloadVariant: &pb.MeshPacket_Decoded{
											Decoded: &pb.Data{
												Portnum:   pb.PortNum_ADMIN_APP,
												Payload:   respBytes,
												RequestId: payload.Packet.Id,
											},
										},
									},
								},
							}); err != nil {
								return fmt.Errorf("writing to streamConn: %w", err)
							}
						}
					}
				}
			}
		}
	})
	// Handle sending messages to client
	eg.Go(func() error {
		ch := make(chan *pb.FromRadio)
		r.mu.Lock()
		r.fromRadioSubscribers[ch] = struct{}{}
		r.mu.Unlock()
		defer func() {
			r.mu.Lock()
			delete(r.fromRadioSubscribers, ch)
			r.mu.Unlock()
		}()

		for {
			select {
			case <-egCtx.Done():
				return nil
			case msg := <-ch:
				if err := streamConn.Write(msg); err != nil {
					return fmt.Errorf("writing to streamConn: %w", err)
				}
			}
		}
	})

	return eg.Wait()
}

func (r *Radio) listenTCP(ctx context.Context) error {
	l, err := net.Listen("tcp", r.cfg.TCPListenAddr)
	if err != nil {
		return fmt.Errorf("listening: %w", err)
	}
	r.logger.Info("listening for tcp connections", "addr", r.cfg.TCPListenAddr)

	for {
		c, err := l.Accept()
		if err != nil {
			r.logger.Error("failed to accept connection", "err", err)
			continue
		}
		go func() {
			if err := r.handleConn(ctx, c); err != nil {
				r.logger.Error("failed to handle TCP connection", "err", err)
			}
		}()
	}
}

// Conn returns an in-memory connection to the emulated radio.
func (r *Radio) Conn(ctx context.Context) net.Conn {
	clientConn, radioConn := net.Pipe()
	go func() {
		if err := r.handleConn(ctx, radioConn); err != nil {
			r.logger.Error("failed to handle in-memory connection", "err", err)
		}
	}()
	return clientConn
}
