package testdata

import (
	contractutils "github.com/xos-labs/node/contracts/utils"
	evmtypes "github.com/xos-labs/node/x/vm/types"
)

func LoadCounterContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("Counter.json")
}

func LoadCounterFactoryContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("CounterFactory.json")
}
