package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getPingBytes() []byte {
	raw_bytes := "45 00 00 54 48 a5 00 00 40 01 1d e5 0a 01 00 0a 0a 01 00 14 08 00 f4 e0 04 20 00 05 5d a2 dd 45 00 04 d9 0a 08 09 0a 0b 0c 0d 0e 0f 10 11 12 13 14 15 16 17 18 19 1a 1b 1c 1d 1e 1f 20 21 22 23 24 25 26 27 28 29 2a 2b 2c 2d 2e 2f 30 31 32 33 34 35 36 37"
	split := strings.Join(strings.Split(raw_bytes, " "), "")
	a, _ := hex.DecodeString(split)
	return a
}

func getTCPBytes() []byte {
	raw_bytes := "45 00 00 40 00 00 40 00 40 06 26 99 0a 01 00 0a 0a 01 00 14 c5 9a 00 50 54 7c 13 59 00 00 00 00 b0 c2 ff ff 31 4f 00 00 02 04 05 b4 01 03 03 06 01 01 08 0a 1e b7 a4 56 00 00 00 00 04 02 00 00"
	split := strings.Join(strings.Split(raw_bytes, " "), "")
	a, _ := hex.DecodeString(split)
	return a
}

func TestPingParse(t *testing.T) {
	bytes := getPingBytes()
	parse(bytes)
}

func TestGetVersion(t *testing.T) {
	assert.Equal(t, uint(4), getVersion(getPingBytes()))
	assert.Equal(t, uint(4), getVersion(getTCPBytes()))
}

func TestGetTCPProtocol(t *testing.T) {
	assert.Equal(t, TCP, getProtocol(getTCPBytes()))
}

func TestGetPingProtocol(t *testing.T)  {
	assert.Equal(t, Ping, getProtocol(getPingBytes()))
}
