package rng

import sdk "github.com/cosmos/cosmos-sdk/types"

type Rng interface {
	GetRandomness(ctx sdk.Context, msg string) []byte
}

type RngProxy struct {
	provablyFair *ProvablyFairRNG
	cosigClient  *CoSigClient
}

func NewRngProxy() *RngProxy {
	return &RngProxy{
		provablyFair: NewProvablyFairRNG(NewRNGConfig([]byte("clientSeed"), []byte("serverSeed"), 0)),
		cosigClient:  NewCoSigClient(),
	}
}

func (r *RngProxy) GetRandomness(ctx sdk.Context, msg string) []byte {
	return r.provablyFair.Next32Bytes()
	//return r.cosigClient.GetRandomness(ctx, msg)
}
