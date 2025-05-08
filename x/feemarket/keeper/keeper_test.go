package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xos-labs/node/testutil/integration/os/network"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestSetGetBlockGasWanted(t *testing.T) {
	var (
		nw  *network.UnitTestNetwork
		ctx sdk.Context
	)
	testCases := []struct {
		name     string
		malleate func()
		expGas   uint64
	}{
		{
			"with last block given",
			func() {
				nw.App.FeeMarketKeeper.SetBlockGasWanted(ctx, uint64(1000000))
			},
			uint64(1000000),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// reset network and context
			nw = network.NewUnitTestNetwork()
			ctx = nw.GetContext()

			tc.malleate()

			gas := nw.App.FeeMarketKeeper.GetBlockGasWanted(ctx)
			require.Equal(t, tc.expGas, gas, tc.name)
		})
	}
}

func TestSetGetGasFee(t *testing.T) {
	var (
		nw  *network.UnitTestNetwork
		ctx sdk.Context
	)
	testCases := []struct {
		name     string
		malleate func()
		expFee   math.LegacyDec
	}{
		{
			"with last block given",
			func() {
				nw.App.FeeMarketKeeper.SetBaseFee(ctx, math.LegacyOneDec())
			},
			math.LegacyOneDec(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// reset network and context
			nw = network.NewUnitTestNetwork()
			ctx = nw.GetContext()

			tc.malleate()

			fee := nw.App.FeeMarketKeeper.GetBaseFee(ctx)
			require.Equal(t, tc.expFee, fee, tc.name)
		})
	}
}
