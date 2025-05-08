package ibc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xos-labs/node/contracts"
	evmibctesting "github.com/xos-labs/node/ibc/testing"
	erc20types "github.com/xos-labs/node/x/erc20/types"
	"github.com/xos-labs/node/xosd"

	ibctesting "github.com/cosmos/ibc-go/v10/testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// NativeErc20Info holds details about a deployed ERC20 token.
type NativeErc20Info struct {
	Denom        string
	ContractAbi  abi.ABI
	ContractAddr common.Address
	Account      common.Address // The address of the minter on the EVM chain
	InitialBal   *big.Int
}

// SetupNativeErc20 deploys, registers, and mints a native ERC20 token on an EVM-based chain.
func SetupNativeErc20(t *testing.T, chain *evmibctesting.TestChain) *NativeErc20Info {
	t.Helper()

	evmCtx := chain.GetContext()
	evmApp := chain.App.(*xosd.XOSD)

	// Deploy new ERC20 contract with default metadata
	contractAddr, err := evmApp.Erc20Keeper.DeployERC20Contract(evmCtx, banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: "example", Exponent: 18},
		},
		Name:   "Example",
		Symbol: "Ex",
	})
	if err != nil {
		t.Fatalf("ERC20 deployment failed: %v", err)
	}
	chain.NextBlock()

	// Register the contract
	_, err = evmApp.Erc20Keeper.RegisterERC20(evmCtx, &erc20types.MsgRegisterERC20{
		Authority:      authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		Erc20Addresses: []string{contractAddr.Hex()},
	})
	if err != nil {
		t.Fatalf("RegisterERC20 failed: %v", err)
	}

	// Mint tokens to default sender
	contractAbi := contracts.ERC20MinterBurnerDecimalsContract.ABI
	nativeDenom := erc20types.CreateDenom(contractAddr.String())
	sendAmt := ibctesting.DefaultCoinAmount
	senderAcc := chain.SenderAccount.GetAddress()

	_, err = evmApp.EVMKeeper.CallEVM(
		evmCtx,
		contractAbi,
		erc20types.ModuleAddress,
		contractAddr,
		true,
		"mint",
		common.BytesToAddress(senderAcc),
		big.NewInt(sendAmt.Int64()),
	)
	if err != nil {
		t.Fatalf("mint call failed: %v", err)
	}

	// Verify minted balance
	bal := evmApp.Erc20Keeper.BalanceOf(evmCtx, contractAbi, contractAddr, common.BytesToAddress(senderAcc))
	if bal.Cmp(big.NewInt(sendAmt.Int64())) != 0 {
		t.Fatalf("unexpected ERC20 balance; got %s, want %s", bal.String(), sendAmt.String())
	}

	return &NativeErc20Info{
		Denom:        nativeDenom,
		ContractAbi:  contractAbi,
		ContractAddr: contractAddr,
		Account:      common.BytesToAddress(senderAcc),
		InitialBal:   big.NewInt(sendAmt.Int64()),
	}
}
