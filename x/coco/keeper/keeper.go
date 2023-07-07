package keeper

import (
	"fmt"
	"github.com/dymensionxyz/rollapp/x/rng"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/dymensionxyz/rollapp/x/coco/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		rng        rng.Rng
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	rng rng.Rng,

) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		rng:        rng,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetRngFromExternal(ctx sdk.Context, msg string) []byte {
	return k.rng.GetRandomness(ctx, msg)
}

func (k Keeper) SetRng(ctx sdk.Context, random []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRngKey(ctx.BlockHeight())
	store.Set(key, random)
}

func (k Keeper) GetRng(ctx sdk.Context, blockHeight int64) []byte {
	store := ctx.KVStore(k.storeKey)
	key := types.GetRngKey(blockHeight)
	return store.Get(key)
}
