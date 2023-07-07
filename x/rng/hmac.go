package rng

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

type RNGConfig struct {
	clientSeed []byte
	serverSeed []byte
	nonce      uint64
}

func NewRNGConfig(clientSeed, serverSeed []byte, nonce uint64) *RNGConfig {
	return &RNGConfig{
		clientSeed: clientSeed,
		serverSeed: serverSeed,
		nonce:      nonce,
	}
}

type ProvablyFairRNG struct {
	config *RNGConfig

	currentRound    int64
	currentRoundMac []byte
}

func NewProvablyFairRNG(config *RNGConfig) *ProvablyFairRNG {
	return &ProvablyFairRNG{
		config:          config,
		currentRound:    0,
		currentRoundMac: nil,
	}
}

func (p *ProvablyFairRNG) updateCurrentRoundBuffer() {
	key := p.config.serverSeed

	input := fmt.Sprintf("%s:%d:%d", p.config.clientSeed, p.config.nonce, p.currentRound)

	mac := HmacSha256(key, []byte(input))

	p.currentRoundMac = mac
}

func (p *ProvablyFairRNG) Next32Bytes() []byte {
	p.updateCurrentRoundBuffer()
	p.currentRound += 1
	return p.currentRoundMac
}

func HmacSha256(key, input []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(input)
	return h.Sum(nil)
}
