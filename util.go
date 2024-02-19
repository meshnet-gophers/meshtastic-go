package meshtastic

import (
	pbuf "github.com/meshnet-gophers/meshtastic-go/meshtastic"
)

type Node struct {
	LongName      string
	ShortName     string
	ID            uint32
	HardwareModel pbuf.HardwareModel
}

// Not actually in use yet 😅
func (n *Node) EncryptPacket(pkt *pbuf.MeshPacket, channelName string, key []byte) *pbuf.MeshPacket {
	payload := pkt.GetPayloadVariant()
	_ = payload
	switch p := payload.(type) {
	case *pbuf.MeshPacket_Decoded:
		_ = p
		encrypted := pbuf.MeshPacket_Encrypted{
			Encrypted: nil,
		}
		_ = encrypted
	}
	return nil
}
