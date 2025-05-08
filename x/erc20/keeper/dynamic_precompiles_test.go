package keeper_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	utiltx "github.com/xos-labs/node/testutil/tx"
	"github.com/xos-labs/node/x/erc20/types"
	"github.com/xos-labs/node/x/vm/statedb"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestRegisterERC20CodeHash() {
	var (
		ctx sdk.Context
		// bytecode and codeHash is the same for all IBC coins
		// cause they're all using the same contract
		bytecode             = common.FromHex(types.Erc20Bytecode)
		codeHash             = crypto.Keccak256(bytecode)
		nonce         uint64 = 10
		balance              = big.NewInt(100)
		emptyCodeHash        = crypto.Keccak256(nil)
	)

	account := utiltx.GenerateAddress()

	testCases := []struct {
		name     string
		malleate func()
		existent bool
	}{
		{
			"ok",
			func() {
			},
			false,
		},
		{
			"existent account",
			func() {
				err := suite.network.App.EVMKeeper.SetAccount(ctx, account, statedb.Account{
					CodeHash: codeHash,
					Nonce:    nonce,
					Balance:  balance,
				})
				suite.Require().NoError(err)
			},
			true,
		},
	}
	for _, tc := range testCases {
		suite.SetupTest() // reset
		ctx = suite.network.GetContext()
		tc.malleate()

		err := suite.network.App.Erc20Keeper.RegisterERC20CodeHash(ctx, account)
		suite.Require().NoError(err)

		acc := suite.network.App.EVMKeeper.GetAccount(ctx, account)
		suite.Require().Equal(codeHash, acc.CodeHash)
		if tc.existent {
			suite.Require().Equal(balance, acc.Balance)
			suite.Require().Equal(nonce, acc.Nonce)
		} else {
			suite.Require().Equal(common.Big0, acc.Balance)
			suite.Require().Equal(uint64(0), acc.Nonce)
		}

		err = suite.network.App.Erc20Keeper.UnRegisterERC20CodeHash(ctx, account)
		suite.Require().NoError(err)

		acc = suite.network.App.EVMKeeper.GetAccount(ctx, account)
		suite.Require().Equal(emptyCodeHash, acc.CodeHash)
		if tc.existent {
			suite.Require().Equal(balance, acc.Balance)
			suite.Require().Equal(nonce, acc.Nonce)
		} else {
			suite.Require().Equal(common.Big0, acc.Balance)
			suite.Require().Equal(uint64(0), acc.Nonce)
		}

	}
}
