package p2p

// Represents remote Node
type Peer interface {
	Close() error
}

// Anything that handles communication between nodes in network
// Can be in form (TCP, UDP, Websockets)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
