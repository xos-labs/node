package constants

import (
	"fmt"

	evmtypes "github.com/xos-labs/node/x/vm/types"
)

const (
	// DefaultGasPrice is used in testing as the default to use for transactions
	DefaultGasPrice = 20

	// ExampleAttoDenom provides an example denom for use in tests
	ExampleAttoDenom = "aatom"

	// ExampleMicroDenom provides an example denom for use in tests
	ExampleMicroDenom = "uatom"

	// XOSDisplayDenom provides an example display denom for use in tests
	XOSDisplayDenom = "atom"

	// ExampleBech32Prefix provides an example Bech32 prefix for use in tests
	ExampleBech32Prefix = "cosmos"

	// ExampleEIP155ChainID provides an example EIP-155 chain ID for use in tests
	ExampleEIP155ChainID = 9001

	// WTOKENContractMainnet is the WTOKEN contract address for mainnet
	WTOKENContractMainnet = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517"
	// WTOKENContractTestnet is the WTOKEN contract address for testnet
	WTOKENContractTestnet = "0xcc491f589b45d4a3c679016195b3fb87d7848210"
)

var (
	// AppChainIDPrefix provides a chain ID prefix for EIP-155 that can be used in tests
	AppChainIDPrefix = fmt.Sprintf("cosmos_%d", ExampleEIP155ChainID)

	// AppChainID provides a chain ID that can be used in tests
	AppChainID = AppChainIDPrefix + "-1"

	// SixDecimalsChainID provides a chain ID which is being set up with 6 decimals
	SixDecimalsChainID = "ossix_9002-2"

	// AppChainCoinInfo provides the coin info for the example chain
	//
	// It is a map of the chain id and its corresponding EvmCoinInfo
	// that allows initializing the app with different coin info based on the
	// chain id
	AppChainCoinInfo = map[string]evmtypes.EvmCoinInfo{
		AppChainID: {
			Denom:        ExampleAttoDenom,
			DisplayDenom: XOSDisplayDenom,
			Decimals:     evmtypes.EighteenDecimals,
		},
		SixDecimalsChainID: {
			Denom:        ExampleMicroDenom,
			DisplayDenom: XOSDisplayDenom,
			Decimals:     evmtypes.SixDecimals,
		},
	}
)
