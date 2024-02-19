package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/meshnet-gophers/meshtastic-go"
	"github.com/meshnet-gophers/meshtastic-go/emulated"
	pb "github.com/meshnet-gophers/meshtastic-go/meshtastic"
	"github.com/meshnet-gophers/meshtastic-go/mqtt"
	"github.com/meshnet-gophers/meshtastic-go/radio"
	"github.com/meshnet-gophers/meshtastic-go/transport"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	// TODO: Flesh this example out and make it configurable
	ctx := context.Background()
	log.SetLevel(log.DebugLevel)

	nodeID, err := meshtastic.RandomNodeID()
	if err != nil {
		panic(err)
	}
	r, err := emulated.NewRadio(emulated.Config{
		LongName:   "EXAMPLE",
		ShortName:  "EMPL",
		NodeID:     nodeID,
		MQTTClient: &mqtt.DefaultClient,
		Channels: &pb.ChannelSet{
			Settings: []*pb.ChannelSettings{
				{
					Name: "LongFast",
					Psk:  radio.DefaultKey,
				},
			},
		},
		BroadcastNodeInfoInterval: 5 * time.Minute,

		BroadcastPositionInterval: 5 * time.Minute,
		// Hardcoded to the position of Buckingham Palace.
		PositionLatitudeI:  515014760,
		PositionLongitudeI: -1406340,
		PositionAltitude:   2,

		TCPListenAddr: "127.0.0.1:4403",
	})
	if err != nil {
		panic(err)
	}

	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		if err := r.Run(egCtx); err != nil {
			return fmt.Errorf("running radio: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		conn, err := transport.NewClientStreamConn(r.Conn(egCtx))
		if err != nil {
			return fmt.Errorf("creating connection: %w", err)
		}

		msg := &pb.ToRadio{
			PayloadVariant: &pb.ToRadio_Packet{
				Packet: &pb.MeshPacket{
					From: nodeID.Uint32(),
					// This is hard coded to Noah's node ID
					To: 2437877602,
					PayloadVariant: &pb.MeshPacket_Decoded{
						Decoded: &pb.Data{
							Portnum: pb.PortNum_TEXT_MESSAGE_APP,
							Payload: []byte("from main!!"),
						},
					},
				},
			},
		}
		if err := conn.Write(msg); err != nil {
			return fmt.Errorf("writing to radio: %w", err)
		}

		for {
			select {
			case <-egCtx.Done():
				return nil
			default:
			}
			msg := &pb.FromRadio{}
			if err := conn.Read(msg); err != nil {
				return fmt.Errorf("reading from radio: %w", err)
			}
			log.Info("FromRadio!!", "packet", msg)
		}
	})

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
