package main

import (
	"github.com/songgao/water"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"tcp-go/main/tcp"
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
	packet := gopacket.NewPacket(bytes, layers.LayerTypeIPv4, gopacket.Default)

	// Get the TCP layer from this packet
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		layer, _ := tcpLayer.(*layers.TCP)
		tcp.Process(*layer)
	} else {
		log.Printf("[ignoring packet] %v\n", packet)
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
