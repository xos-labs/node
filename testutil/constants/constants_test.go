package constants_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	chainconfig "github.com/xos-labs/node/cmd/xosd/config"
	"github.com/xos-labs/node/testutil/constants"
	"github.com/xos-labs/node/xosd"
)

func TestRequireSameTestDenom(t *testing.T) {
	require.Equal(t,
		constants.ExampleAttoDenom,
		xosd.XOSChainDenom,
		"test denoms should be the same across the repo",
	)
}

func TestRequireSameTestBech32Prefix(t *testing.T) {
	require.Equal(t,
		constants.ExampleBech32Prefix,
		chainconfig.Bech32Prefix,
		"bech32 prefixes should be the same across the repo",
	)
}

func TestRequireSameWTOKENMainnet(t *testing.T) {
	require.Equal(t,
		constants.WTOKENContractMainnet,
		xosd.WTOKENContractMainnet,
		"wtoken contract addresses should be the same across the repo",
	)
}
