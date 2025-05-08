package network

import (
	testconstants "github.com/xos-labs/node/testutil/constants"
)

// chainsWTOKENHex is an utility map used to retrieve the WTOKEN contract
// address in hex format from the chain ID.
//
// TODO: refactor to define this in the example chain initialization and pass as function argument
var chainsWTOKENHex = map[string]string{
	testconstants.AppChainID: testconstants.WTOKENContractMainnet,
}

// GetWTOKENContractHex returns the hex format of address for the WTOKEN contract
// given the chainID. If the chainID is not found, it defaults to the mainnet
// address.
func GetWTOKENContractHex(chainID string) string {
	address, found := chainsWTOKENHex[chainID]

	// default to mainnet address
	if !found {
		address = chainsWTOKENHex[testconstants.AppChainID]
	}

	return address
}
