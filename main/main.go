package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/songgao/water"
	"log"
	"tcp-go/main/tcp"
	"encoding/hex"
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
		parse(packet[:n], ifce)
	}
}

func parse(bytes []byte, ifce *water.Interface) {
	// Decode a packet
	packet := gopacket.NewPacket(bytes, layers.LayerTypeIPv4, gopacket.Default)

	layer := getLayer(packet, layers.LayerTypeTCP)
	tcpLayer, _ := layer.(*layers.TCP)
	ipv4Layer := getLayer(packet, layers.LayerTypeIPv4).(*layers.IPv4)

	if tcpLayer == nil {
		log.Printf("[ignoring packet] %v\n", packet)
		return
	}

	tcpResult := tcp.Process(*tcpLayer)
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths: false,
		ComputeChecksums: false,
	}

	err := gopacket.SerializeLayers(buf, opts,
		&layers.IPv4{
			Version: ipv4Layer.Version,
			Protocol: ipv4Layer.Protocol,
			IHL: 5,
			SrcIP: ipv4Layer.DstIP,
			DstIP: ipv4Layer.SrcIP,
		},
		tcpResult,
		gopacket.Payload([]byte{}))

	if err != nil {
		log.Printf("[error serializing buffer] layer: %v, err: %v", layer, err)
		return
	}
	if buf != nil {
		sze, err := ifce.Write(buf.Bytes())
		log.Printf("[write success] sze: %v, err: %v, buf: %v", sze, err, hex.EncodeToString(buf.Bytes()))
	}
}

func getLayer(packet gopacket.Packet, layerType gopacket.LayerType) gopacket.Layer {
	layer := packet.Layer(layerType)
	if layer == nil {
		log.Printf("[ignoring packet] %v\n", packet)
		return nil
	}
	return layer
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
