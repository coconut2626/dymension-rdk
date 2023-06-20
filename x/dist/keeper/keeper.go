package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/dymensionxyz/rollapp/x/dist/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	distkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	distkeeper.Keeper

	authKeeper    disttypes.AccountKeeper
	bankKeeper    disttypes.BankKeeper
	stakingKeeper types.StakingKeeper
	seqKeeper     types.SequencerKeeper

	paramSpace paramtypes.Subspace

	blockedAddrs     map[string]bool
	feeCollectorName string
}

// NewKeeper creates a new distribution Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey, paramSpace paramtypes.Subspace,
	ak disttypes.AccountKeeper, bk disttypes.BankKeeper, sk types.StakingKeeper, seqk types.SequencerKeeper,
	feeCollectorName string, blockedAddrs map[string]bool,
) Keeper {
	k := distkeeper.NewKeeper(cdc, key, paramSpace, ak, bk, sk, feeCollectorName)
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		Keeper:           k,
		authKeeper:       ak,
		bankKeeper:       bk,
		stakingKeeper:    sk,
		seqKeeper:        seqk,
		paramSpace:       paramSpace,
		blockedAddrs:     blockedAddrs,
		feeCollectorName: feeCollectorName,
	}
}
