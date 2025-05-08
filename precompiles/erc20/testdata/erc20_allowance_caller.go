package testdata

import (
	contractutils "github.com/xos-labs/node/contracts/utils"
	evmtypes "github.com/xos-labs/node/x/vm/types"
)

func LoadERC20AllowanceCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20AllowanceCaller.json")
}
