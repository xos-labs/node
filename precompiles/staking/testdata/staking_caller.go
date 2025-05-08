package testdata

import (
	contractutils "github.com/xos-labs/node/contracts/utils"
	evmtypes "github.com/xos-labs/node/x/vm/types"
)

func LoadStakingCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("StakingCaller.json")
}
