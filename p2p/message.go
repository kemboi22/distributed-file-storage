package p2p

import "net"

// Message represents any arbitiary data that is being sent over each transport
// betweem 2 nodes in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
