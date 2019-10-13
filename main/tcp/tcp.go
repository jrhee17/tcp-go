package tcp

import (
	"log"
	"github.com/google/gopacket/layers"
)

func Process(layer layers.TCP) {
	log.Println(layer)
}
