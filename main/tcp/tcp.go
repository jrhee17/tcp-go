package tcp

import (
	"github.com/google/gopacket/layers"
	"log"
)

func Process(layer layers.TCP) *layers.TCP {
	if (layer.SYN == false) {
		log.Printf("[ignoring non syn] layer: %v", layer)
		return nil
	}

	log.Printf("[processing] layer: %v", layer)

	tcpLayer := layers.TCP{
		ACK: true,
		SYN: true,
	}

	return &tcpLayer
}
