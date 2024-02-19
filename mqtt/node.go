package mqtt

import meshtastic "github.com/meshnet-gophers/meshtastic-go/meshtastic"

// Node implements a meshtastic node that connects only via MQTT
type Node struct {
	user *meshtastic.User
}

func NewNode(user *meshtastic.User) *Node {
	return &Node{
		user: user,
	}
}
