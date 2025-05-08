package testdata

import (
	contractutils "github.com/xos-labs/node/contracts/utils"
	evmtypes "github.com/xos-labs/node/x/vm/types"
)

func LoadWTOKEN9TestCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("WEVMOS9TestCaller.json")
}
