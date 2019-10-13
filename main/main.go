package main

import (
	"github.com/songgao/water"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"fmt"
)

func main() {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", ifce.Name())

	packet := make([]byte, 2000)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		parse(packet[:n])
	}
}

func parse(bytes []byte) {
	// Decode a packet
	packet := gopacket.NewPacket(bytes, layers.LayerTypeEthernet, gopacket.Default)
	log.Printf("[parse] packet: %v\n", packet.LinkLayer())
	// Get the TCP layer from this packet
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		fmt.Println("This is a TCP packet!")
		// Get actual TCP data from this layer
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("From src port %d to dst port %d\n", tcp.SrcPort, tcp.DstPort)
	}
}

func getVersion(bytes []byte) uint {
	return uint(bytes[0] >> 4)
}

func getProtocol(bytes []byte) Protocol {
	switch bytes[9] {
	case 1:
		return Ping
	case 6:
		return TCP
	default:
		return Unknown
	}
}

type Protocol int

const (
	Ping Protocol = 1
	TCP Protocol = 6
	Unknown Protocol = -1
)
