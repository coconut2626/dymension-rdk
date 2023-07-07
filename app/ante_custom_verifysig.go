package app

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator will not check signatures on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak              authante.AccountKeeper
	signModeHandler authsigning.SignModeHandler
}

func NewCustomSigVerificationDecorator(ak authante.AccountKeeper, signModeHandler authsigning.SignModeHandler) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:              ak,
		signModeHandler: signModeHandler,
	}
}

// OnlyLegacyAminoSigners checks SignatureData to see if all
// signers are using SIGN_MODE_LEGACY_AMINO_JSON. If this is the case
// then the corresponding SignatureV2 struct will not have account sequence
// explicitly set, and we should skip the explicit verification of sig.Sequence
// in the SigVerificationDecorator's AnteHandler function.
func OnlyLegacyAminoSigners(sigData signing.SignatureData) bool {
	switch v := sigData.(type) {
	case *signing.SingleSignatureData:
		return v.SignMode == signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	case *signing.MultiSignatureData:
		for _, s := range v.Signatures {
			if !OnlyLegacyAminoSigners(s) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// stdSigs contains the sequence number, account number, and signatures.
	// When simulating, this would just be a 0-length slice.
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return ctx, err
	}

	signerAddrs := sigTx.GetSigners()

	// check that signer length and signature length are the same
	if len(sigs) != len(signerAddrs) {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "invalid number of signer;  expected: %d, got %d", len(signerAddrs), len(sigs))
	}

	for i, sig := range sigs {
		acc, err := GetSignerAcc(ctx, svd.ak, signerAddrs[i])
		if err != nil {
			return ctx, err
		}

		// retrieve pubkey
		pubKey := acc.GetPubKey()
		if !simulate && pubKey == nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
		}

		if !ctx.IsCheckTx() && !ctx.IsReCheckTx() {
			// Check account sequence number.
			if sig.Sequence != acc.GetSequence() {
				return ctx, sdkerrors.Wrapf(
					sdkerrors.ErrWrongSequence,
					"account sequence mismatch, expected %d, got %d", acc.GetSequence(), sig.Sequence,
				)
			}
		}

		// retrieve signer data
		genesis := ctx.BlockHeight() == 0
		chainID := ctx.ChainID()
		var accNum uint64
		if !genesis {
			accNum = acc.GetAccountNumber()
		}
		signerData := authsigning.SignerData{
			Address:       acc.GetAddress().String(),
			ChainID:       chainID,
			AccountNumber: accNum,
			Sequence:      acc.GetSequence(),
			PubKey:        pubKey,
		}

		// no need to verify signatures on recheck tx
		if !simulate && !ctx.IsReCheckTx() {
			err := authsigning.VerifySignature(pubKey, signerData, sig.Data, svd.signModeHandler, tx)
			if err != nil {
				var errMsg string
				if OnlyLegacyAminoSigners(sig.Data) {
					// If all signers are using SIGN_MODE_LEGACY_AMINO, we rely on VerifySignature to check account sequence number,
					// and therefore communicate sequence number as a potential cause of error.
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d), sequence (%d) and chain-id (%s)", accNum, acc.GetSequence(), chainID)
				} else {
					errMsg = fmt.Sprintf("signature verification failed; please verify account number (%d) and chain-id (%s)", accNum, chainID)
				}
				return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, errMsg)

			}
		}
	}

	return next(ctx, tx, simulate)
}

// IncrementSequenceDecorator handles incrementing sequences of all signers.
// Use the IncrementSequenceDecorator decorator to prevent replay attacks. Note,
// there is need to execute IncrementSequenceDecorator on RecheckTx since
// BaseApp.Commit() will set the check state based on the latest header.
//
// NOTE: Since CheckTx and DeliverTx state are managed separately, subsequent and
// sequential txs orginating from the same account cannot be handled correctly in
// a reliable way unless sequence numbers are managed and tracked manually by a
// client. It is recommended to instead use multiple messages in a tx.
type IncrementSequenceDecorator struct {
	ak authante.AccountKeeper
}

func NewCustomIncrementSequenceDecorator(ak authante.AccountKeeper) IncrementSequenceDecorator {
	return IncrementSequenceDecorator{
		ak: ak,
	}
}

func (isd IncrementSequenceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if ctx.IsCheckTx() || ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}

	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// increment sequence of all signers
	for _, addr := range sigTx.GetSigners() {
		acc := isd.ak.GetAccount(ctx, addr)
		if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
			panic(err)
		}

		isd.ak.SetAccount(ctx, acc)
	}

	return next(ctx, tx, simulate)
}

// GetSignerAcc returns an account for a given address that is expected to sign
// a transaction.
func GetSignerAcc(ctx sdk.Context, ak authante.AccountKeeper, addr sdk.AccAddress) (types.AccountI, error) {
	if acc := ak.GetAccount(ctx, addr); acc != nil {
		return acc, nil
	}

	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", addr)
}
