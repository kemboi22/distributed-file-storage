package p2p

// Represents remote Node
type Peer interface{}

// Anything that handles communication between nodes in network
// Can be in form (TCP, UDP, Websockets)
type Transport interface {
	ListenAndAccept() error
}
