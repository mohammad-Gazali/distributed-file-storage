package main

import (
	"log"

	p2p "github.com/mohammad-Gazali/distributed-file-storage/peer_to_peer"
)

func main() {
	tcpOpts :=  p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}