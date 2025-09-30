package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kemboi22/distributed-file-storage/p2p"
)

func OnPeer(p2p.Peer) error {
	fmt.Printf("Doing Some logical stuff here")
	return nil
}

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)
	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}
	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
