package p2p

import (
	"fmt"
	"log"
	"net"
)

// TCPPeer represents the remote node iver a TCP established connection
type TCPPeer struct {
	// conn is the under lying connection of the peer
	conn net.Conn

	// if we dial and retrive a connection => outbound == true
	// if we dial and retrive a connection => outbound == false
	outbound bool
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	Listener net.Listener
	rpcch    chan RPC
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume() impliments the transport interface will return a read only channel
// for reading incoming messages received from another peer
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) Close() error {
	return t.Listener.Close()
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.Listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	log.Printf("tcp transport listening on port: %s\n", t.ListenAddr)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s", err)
		}
		fmt.Printf("new incoming connection %+v \n", conn)
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("dropping peer connection: %s", err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("TCP handshake error: %s \n", err)
		conn.Close()
		return
	}
	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			return
		}
	}
	rpc := &RPC{}
	for {
		err := t.Decoder.Decode(conn, rpc)
		if err != nil {
			fmt.Printf("tcp read error: %s \n", err)
			return
		}
		rpc.From = conn.RemoteAddr()
		fmt.Printf("message: %+v \n", rpc)
		t.rpcch <- *rpc
	}
}
