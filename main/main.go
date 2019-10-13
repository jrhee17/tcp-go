package main

import (
	"github.com/songgao/water"
	"log"
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
	log.Printf("[Packet Received] Len: %d, Bytes: % x\n", len(bytes), bytes)
	version := getVersion(bytes)
	log.Printf("[parse] version: %d\n", version)
	protocol := getProtocol(bytes)
	log.Printf("[parse] protocol: %d\n", protocol)
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
