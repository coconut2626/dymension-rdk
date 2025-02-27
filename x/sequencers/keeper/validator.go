package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/dymensionxyz/rollapp/x/sequencers/types"
)

//TODO: rename all to sequencer

/* -------------------------------------------------------------------------- */
/*                                    Alias func                              */
/* -------------------------------------------------------------------------- */
// Validator gets the Validator interface for a particular address
func (k Keeper) Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI {
	val, found := k.GetValidator(ctx, address)
	if !found {
		return nil
	}

	return val
}

// ValidatorByConsAddr gets the validator interface for a particular pubkey
func (k Keeper) ValidatorByConsAddr(ctx sdk.Context, addr sdk.ConsAddress) stakingtypes.ValidatorI {
	val, found := k.GetValidatorByConsAddr(ctx, addr)
	if !found {
		return nil
	}

	return val
}

/* -------------------------------------------------------------------------- */
/*                               implementation                              */
/* -------------------------------------------------------------------------- */
/* --------------------------------- GETTERS -------------------------------- */
// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}

	validator = stakingtypes.MustUnmarshalValidator(k.cdc, value)
	return validator, true
}

// get a single validator by consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)
	opAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if opAddr == nil {
		return validator, false
	}

	return k.GetValidator(ctx, opAddr)
}

// get a single sequencer registered on dymint by consensus address
func (k Keeper) GetDymintSequencerByAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (power uint64, found bool) {
	store := ctx.KVStore(k.storeKey)
	powerByte := store.Get(types.GetDymintSeqKey(consAddr))
	if powerByte == nil {
		return 0, false
	}

	return binary.LittleEndian.Uint64(powerByte), true
}

/* --------------------------------- SETTERS -------------------------------- */
// set the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator stakingtypes.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := stakingtypes.MustMarshalValidator(k.cdc, &validator)
	store.Set(types.GetValidatorKey(validator.GetOperator()), bz)
}

// validator index
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator stakingtypes.Validator) error {
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorByConsAddrKey(consAddr), validator.GetOperator())

	return nil
}

// get a single sequencer registered on dymint by consensus address
func (k Keeper) SetDymintSequencerByAddr(ctx sdk.Context, validator stakingtypes.Validator) error {
	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}

	power := validator.ConsensusPower(sdk.DefaultPowerReduction)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(power))

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetDymintSeqKey(consAddr), b)
	return nil
}

// get groups of validators

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validator := stakingtypes.MustUnmarshalValidator(k.cdc, iterator.Value())
		validators = append(validators, validator)
	}

	return validators
}
