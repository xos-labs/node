package network

import (
	erc20types "github.com/xos-labs/node/x/erc20/types"
	feemarkettypes "github.com/xos-labs/node/x/feemarket/types"
	evmtypes "github.com/xos-labs/node/x/vm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func getQueryHelper(ctx sdktypes.Context, encCfg testutil.TestEncodingConfig) *baseapp.QueryServiceTestHelper {
	interfaceRegistry := encCfg.InterfaceRegistry
	// This is needed so that state changes are not committed in precompiles
	// simulations.
	cacheCtx, _ := ctx.CacheContext()
	return baseapp.NewQueryServerTestHelper(cacheCtx, interfaceRegistry)
}

func (n *IntegrationNetwork) GetERC20Client() erc20types.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	erc20types.RegisterQueryServer(queryHelper, n.app.Erc20Keeper)
	return erc20types.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetEvmClient() evmtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	evmtypes.RegisterQueryServer(queryHelper, n.app.EVMKeeper)
	return evmtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetGovClient() govtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	govtypes.RegisterQueryServer(queryHelper, govkeeper.NewQueryServer(&n.app.GovKeeper))
	return govtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetBankClient() banktypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	banktypes.RegisterQueryServer(queryHelper, n.app.BankKeeper)
	return banktypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetFeeMarketClient() feemarkettypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	feemarkettypes.RegisterQueryServer(queryHelper, n.app.FeeMarketKeeper)
	return feemarkettypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetAuthClient() authtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	authtypes.RegisterQueryServer(queryHelper, authkeeper.NewQueryServer(n.app.AccountKeeper))
	return authtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetAuthzClient() authz.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	authz.RegisterQueryServer(queryHelper, n.app.AuthzKeeper)
	return authz.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetStakingClient() stakingtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	stakingtypes.RegisterQueryServer(queryHelper, stakingkeeper.Querier{Keeper: n.app.StakingKeeper})
	return stakingtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetDistrClient() distrtypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	distrtypes.RegisterQueryServer(queryHelper, distrkeeper.Querier{Keeper: n.app.DistrKeeper})
	return distrtypes.NewQueryClient(queryHelper)
}

func (n *IntegrationNetwork) GetMintClient() minttypes.QueryClient {
	queryHelper := getQueryHelper(n.GetContext(), n.GetEncodingConfig())
	minttypes.RegisterQueryServer(queryHelper, mintkeeper.NewQueryServerImpl(n.app.MintKeeper))
	return minttypes.NewQueryClient(queryHelper)
}
