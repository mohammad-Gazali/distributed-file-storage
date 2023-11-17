package peertopeer

import "net"

// RPC holds any arbitrary data that is being sent over
// each transport between two nodes in the network.
type RPC struct {
	From    net.Addr
	Payload []byte
}
