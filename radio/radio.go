package radio

import (
	"encoding/base64"
	"errors"
	"fmt"
	generated "github.com/meshnet-gophers/meshtastic-go/meshtastic"
	"google.golang.org/protobuf/proto"
)

// not sure what i was getting at with this
type comMode uint8

const (
	ComModeProtobuf comMode = iota + 1
	ComModeSerialDebug
)

// Transport defines methods required for communicating with a radio via serial, ble, or tcp
// Probably need to reevaluate this to just use the ToRadio and FromRadio protobufs
type Transport interface {
	Connect() error
	SendPacket(data []byte) error
	RequestConfig() error

	//	Listen(ch chan)
	Close() error
}

// Something is something created to track keys for packet decrypting
type Something struct {
	keys map[string][]byte
}

// default encryption key, commonly referenced as AQ==
// as base64: 1PG7OiApB1nwvP+rz05pAQ==
var DefaultKey = []byte{0xd4, 0xf1, 0xbb, 0x3a, 0x20, 0x29, 0x07, 0x59, 0xf0, 0xbc, 0xff, 0xab, 0xcf, 0x4e, 0x69, 0x01}

// ParseKey converts the most common representation of a channel encryption key (URL encoded base64) to a byte slice
func ParseKey(key string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(key)
}

func NewThing() *Something {
	return &Something{keys: map[string][]byte{
		"LongFast":  DefaultKey,
		"LongSlow":  DefaultKey,
		"VLongSlow": DefaultKey,
	}}
}

// GenerateByteSlices creates a bunch of weak keys for use when interfacing on MQTT.
// This creates 128, 192, and 256 bit AES keys with only a single byte specified
func GenerateByteSlices() [][]byte {
	// There are 256 possible values for a single byte
	// We create 1536 slices: 512 with 16 bytes, 512 with 24 bytes, and 512 with 32 bytes
	allSlices := make([][]byte, 256*3)

	for i := 0; i < 256; i++ {
		// Create a slice of 16 bytes for the first 256 slices
		slice16 := make([]byte, 16)
		// Set the last byte to the current iteration value.
		slice16[15] = byte(i)
		// Assign the slice to our slice of slices.
		allSlices[i] = slice16

		// Create a slice of 24 bytes (192 bits) for the next 256 slices
		slice24 := make([]byte, 24)
		// Set the last byte to the current iteration value.
		slice24[23] = byte(i)
		// Assign the slice to our slice of slices, offset by 256.
		allSlices[i+256] = slice24

		// Create a slice of 32 bytes for the last 256 slices
		slice32 := make([]byte, 32)
		// Set the last byte to the current iteration value.
		slice32[31] = byte(i)
		// Assign the slice to our slice of slices, offset by 512.
		allSlices[i+512] = slice32
	}

	return allSlices
}

var ErrDecrypt = errors.New("unable to decrypt payload")

// xorHash computes a simple XOR hash of the provided byte slice.
func xorHash(p []byte) uint8 {
	var code uint8
	for _, b := range p {
		code ^= b
	}
	return code
}

// GenerateHash returns the hash for a given channel by XORing the channel name and PSK.
func ChannelHash(channelName string, channelKey []byte) (uint32, error) {
	if len(channelKey) == 0 {
		return 0, fmt.Errorf("channel key cannot be empty")
	}

	h := xorHash([]byte(channelName))
	h ^= xorHash(channelKey)

	return uint32(h), nil
}

// Attempts to decrypt a packet with the specified key, or return the already decrypted data if present.
func TryDecode(packet *generated.MeshPacket, key []byte) (*generated.Data, error) {
	//packet := env.GetPacket()
	switch packet.GetPayloadVariant().(type) {
	case *generated.MeshPacket_Decoded:
		//fmt.Println("decoded")
		return packet.GetDecoded(), nil
	case *generated.MeshPacket_Encrypted:
		//	fmt.Println("encrypted")
		/*
			key, exists := s.keys[env.ChannelId]
			if !exists {
				return nil, errors.New("no decryption key for channel")
			}

		*/
		decrypted, err := XOR(packet.GetEncrypted(), key, packet.Id, packet.From)
		if err != nil {
			//log.Error("failed decrypting packet", "error", err)
			return nil, ErrDecrypt
		}

		var meshPacket generated.Data
		err = proto.Unmarshal(decrypted, &meshPacket)
		if err != nil {
			//		log.Info("failed with supplied key")
			return nil, ErrDecrypt
		}
		//fmt.Println("supplied key success")
		return &meshPacket, nil
	default:
		return nil, errors.New("unknown payload variant")
	}
}

// decode a payload to a Data protobuf
func (s *Something) TryDecode(packet *generated.MeshPacket, key []byte) (*generated.Data, error) {
	return TryDecode(packet, key)
}
