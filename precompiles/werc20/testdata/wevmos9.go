package testdata

import (
	contractutils "github.com/xos-labs/node/contracts/utils"
	evmtypes "github.com/xos-labs/node/x/vm/types"
)

// LoadWTOKEN9Contract load the WTOKEN9 contract from the json representation of
// the Solidity contract.
func LoadWTOKEN9Contract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("WTOKEN9.json")
}
