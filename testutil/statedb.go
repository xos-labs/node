package testutil

import (
	"github.com/ethereum/go-ethereum/common"
	anteinterfaces "github.com/xos-labs/node/ante/interfaces"
	"github.com/xos-labs/node/x/vm/statedb"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStateDB returns a new StateDB for testing purposes.
func NewStateDB(ctx sdk.Context, evmKeeper anteinterfaces.EVMKeeper) *statedb.StateDB {
	return statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash())))
}
