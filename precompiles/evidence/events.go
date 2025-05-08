package evidence

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	cmn "github.com/xos-labs/node/precompiles/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// EventTypeSubmitEvidence defines the event type for the evidence SubmitEvidence transaction.
	EventTypeSubmitEvidence = "SubmitEvidence"
)

// EmitSubmitEvidenceEvent creates a new event emitted on a SubmitEvidence transaction.
func (p Precompile) EmitSubmitEvidenceEvent(ctx sdk.Context, stateDB vm.StateDB, origin common.Address, evidenceHash []byte) error {
	// Prepare the event topics
	event := p.ABI.Events[EventTypeSubmitEvidence]
	topics := make([]common.Hash, 2)

	// The first topic is always the signature of the event.
	topics[0] = event.ID

	var err error
	topics[1], err = cmn.MakeTopic(origin)
	if err != nil {
		return err
	}

	// Pack the evidence hash
	arguments := abi.Arguments{event.Inputs[1]}
	packed, err := arguments.Pack(evidenceHash)
	if err != nil {
		return err
	}

	stateDB.AddLog(&ethtypes.Log{
		Address:     p.Address(),
		Topics:      topics,
		Data:        packed,
		BlockNumber: uint64(ctx.BlockHeight()), //nolint:gosec // G115
	})

	return nil
}
