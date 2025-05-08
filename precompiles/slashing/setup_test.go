package slashing_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xos-labs/node/precompiles/slashing"
	"github.com/xos-labs/node/testutil/integration/os/factory"
	"github.com/xos-labs/node/testutil/integration/os/grpc"
	testkeyring "github.com/xos-labs/node/testutil/integration/os/keyring"
	"github.com/xos-labs/node/testutil/integration/os/network"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PrecompileTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile *slashing.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(3)
	var err error
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
		network.WithValidatorOperators([]sdk.AccAddress{
			keyring.GetAccAddr(0),
			keyring.GetAccAddr(1),
			keyring.GetAccAddr(2),
		}),
	)

	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	s.network = nw
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring

	if s.precompile, err = slashing.NewPrecompile(
		s.network.App.SlashingKeeper,
		s.network.App.AuthzKeeper,
	); err != nil {
		panic(err)
	}
}
