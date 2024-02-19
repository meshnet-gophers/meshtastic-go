package transport

import (
	"bytes"
	pb "github.com/meshnet-gophers/meshtastic-go/meshtastic"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
	"net"
	"testing"
)

func TestStreamConn(t *testing.T) {
	radioNetConn, clientNetConn := net.Pipe()
	var client *StreamConn
	var radio *StreamConn

	// Test client -> radio
	sent := &pb.ToRadio{
		PayloadVariant: &pb.ToRadio_WantConfigId{
			WantConfigId: 123,
		},
	}
	received := &pb.ToRadio{}
	eg := errgroup.Group{}
	eg.Go(func() error {
		var err error
		client, err = NewClientStreamConn(clientNetConn)
		require.NoError(t, err)
		return client.Write(sent)
	})
	eg.Go(func() error {
		radio = NewRadioStreamConn(radioNetConn)
		return radio.Read(received)
	})
	require.NoError(t, eg.Wait())
	require.True(t, proto.Equal(sent, received))

	// Test radio -> client
	replySent := &pb.FromRadio{
		Id: 123,
		PayloadVariant: &pb.FromRadio_Config{
			Config: &pb.Config{
				PayloadVariant: &pb.Config_Device{
					Device: &pb.Config_DeviceConfig{
						Role: pb.Config_DeviceConfig_ROUTER,
					},
				},
			},
		},
	}
	replyReceived := &pb.FromRadio{}
	eg = errgroup.Group{}
	eg.Go(func() error {
		return radio.Write(replySent)
	})
	eg.Go(func() error {
		return client.Read(replyReceived)
	})
	require.NoError(t, eg.Wait())
	require.True(t, proto.Equal(replySent, replyReceived))
}

func Test_writeStreamHeader(t *testing.T) {
	out := bytes.NewBuffer(nil)
	err := writeStreamHeader(out, 257)
	require.NoError(t, err)
	require.Equal(t, []byte{Start1, Start2, 0x01, 0x01}, out.Bytes())
}
