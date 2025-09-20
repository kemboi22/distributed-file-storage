package p2p

type Handshaker interface {
	Handshake() error
}

// HandshakeFunc ....?
type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
