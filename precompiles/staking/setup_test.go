package staking_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xos-labs/node/precompiles/staking"
	"github.com/xos-labs/node/testutil/integration/os/factory"
	"github.com/xos-labs/node/testutil/integration/os/grpc"
	testkeyring "github.com/xos-labs/node/testutil/integration/os/keyring"
	"github.com/xos-labs/node/testutil/integration/os/network"
)

type PrecompileTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	bondDenom  string
	precompile *staking.Precompile
}

func TestPrecompileUnitTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)
	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	ctx := nw.GetContext()
	sk := nw.App.StakingKeeper
	bondDenom, err := sk.BondDenom(ctx)
	if err != nil {
		panic(err)
	}

	s.bondDenom = bondDenom
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring
	s.network = nw

	if s.precompile, err = staking.NewPrecompile(
		*s.network.App.StakingKeeper,
		s.network.App.AuthzKeeper,
	); err != nil {
		panic(err)
	}
}
