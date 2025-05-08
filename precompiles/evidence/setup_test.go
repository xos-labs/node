package evidence_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xos-labs/node/precompiles/evidence"
	"github.com/xos-labs/node/testutil/integration/os/factory"
	"github.com/xos-labs/node/testutil/integration/os/grpc"
	testkeyring "github.com/xos-labs/node/testutil/integration/os/keyring"
	"github.com/xos-labs/node/testutil/integration/os/network"

	"cosmossdk.io/x/evidence/exported"
	"cosmossdk.io/x/evidence/types"
)

type PrecompileTestSuite struct {
	suite.Suite

	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile *evidence.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	var err error
	nw := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)

	grpcHandler := grpc.NewIntegrationHandler(nw)
	txFactory := factory.New(nw, grpcHandler)

	router := types.NewRouter()
	router = router.AddRoute(types.RouteEquivocation, testEquivocationHandler(nw.App.EvidenceKeeper))
	nw.App.EvidenceKeeper.SetRouter(router)

	s.network = nw
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring

	if s.precompile, err = evidence.NewPrecompile(
		s.network.App.EvidenceKeeper,
		s.network.App.AuthzKeeper,
	); err != nil {
		panic(err)
	}
}

func testEquivocationHandler(_ interface{}) types.Handler {
	return func(_ context.Context, e exported.Evidence) error {
		if err := e.ValidateBasic(); err != nil {
			return err
		}

		ee, ok := e.(*types.Equivocation)
		if !ok {
			return fmt.Errorf("unexpected evidence type: %T", e)
		}
		if ee.Height%2 == 0 {
			return fmt.Errorf("unexpected even evidence height: %d", ee.Height)
		}

		return nil
	}
}
