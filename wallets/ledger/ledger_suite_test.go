package ledger_test

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/xos-labs/node/testutil/constants"
	"github.com/xos-labs/node/testutil/integration/os/network"
	"github.com/xos-labs/node/wallets/ledger"
	"github.com/xos-labs/node/wallets/ledger/mocks"
	"github.com/xos-labs/node/wallets/usbwallet"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txTypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	auxTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type LedgerTestSuite struct {
	suite.Suite
	txAmino    []byte
	txProtobuf []byte
	ledger     ledger.CosmosEVMSECP256K1
	mockWallet *mocks.Wallet
	hrp        string
}

// Load encoding config for sign doc encoding/decoding
// This is done on app instantiation.
// We use the testutil network to load the encoding config
func init() {
	network.New()
}

func TestLedgerTestSuite(t *testing.T) {
	suite.Run(t, new(LedgerTestSuite))
}

func (suite *LedgerTestSuite) SetupTest() {
	suite.hrp = "cosmos"

	suite.txAmino = suite.getMockTxAmino()
	suite.txProtobuf = suite.getMockTxProtobuf()

	hub, err := usbwallet.NewLedgerHub()
	suite.Require().NoError(err)

	mockWallet := new(mocks.Wallet)
	suite.mockWallet = mockWallet
	suite.ledger = ledger.CosmosEVMSECP256K1{Hub: hub, PrimaryWallet: mockWallet}
}

func (suite *LedgerTestSuite) newPubKey(pk string) (res cryptoTypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	suite.Require().NoError(err)

	pubkey := &ed25519.PubKey{Key: pkBytes}

	return pubkey
}

func (suite *LedgerTestSuite) getMockTxAmino() []byte {
	whitespaceRegex := regexp.MustCompile(`\s+`)
	tmp := whitespaceRegex.ReplaceAllString(fmt.Sprintf(
		`{
			"account_number": "0",
			"chain_id":"%s",
			"fee":{
				"amount":[{"amount":"150","denom":"atom"}],
				"gas":"20000"
			},
			"memo":"memo",
			"msgs":[{
				"type":"cosmos-sdk/MsgSend",
				"value":{
					"amount":[{"amount":"150","denom":"atom"}],
					"from_address":"cosmos10jmp6sgh4cc6zt3e8gw05wavvejgr5pwsjskvv",
					"to_address":"cosmos1fx944mzagwdhx0wz7k9tfztc8g3lkfk6rrgv6l"
				}
			}],
			"sequence":"6"
		}`, constants.AppChainID),
		"",
	)

	return []byte(tmp)
}

func (suite *LedgerTestSuite) getMockTxProtobuf() []byte {
	marshaler := codec.NewProtoCodec(codecTypes.NewInterfaceRegistry())

	memo := "memo"
	msg := bankTypes.NewMsgSend(
		sdk.MustAccAddressFromBech32("cosmos1r5sckdd808qvg7p8d0auaw896zcluqfd7djffp"),
		sdk.MustAccAddressFromBech32("cosmos10t8ca2w09ykd6ph0agdz5stvgau47whhaggl9a"),
		[]sdk.Coin{
			{
				Denom:  "atom",
				Amount: math.NewIntFromUint64(150),
			},
		},
	)

	msgAsAny, err := codecTypes.NewAnyWithValue(msg)
	suite.Require().NoError(err)

	body := &txTypes.TxBody{
		Messages: []*codecTypes.Any{
			msgAsAny,
		},
		Memo: memo,
	}

	pubKey := suite.newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB50")

	pubKeyAsAny, err := codecTypes.NewAnyWithValue(pubKey)
	suite.Require().NoError(err)

	signingMode := txTypes.ModeInfo_Single_{
		Single: &txTypes.ModeInfo_Single{
			Mode: signing.SignMode_SIGN_MODE_DIRECT,
		},
	}

	signerInfo := &txTypes.SignerInfo{
		PublicKey: pubKeyAsAny,
		ModeInfo: &txTypes.ModeInfo{
			Sum: &signingMode,
		},
		Sequence: 6,
	}

	fee := txTypes.Fee{Amount: sdk.NewCoins(sdk.NewInt64Coin("atom", 150)), GasLimit: 20000}

	authInfo := &txTypes.AuthInfo{
		SignerInfos: []*txTypes.SignerInfo{signerInfo},
		Fee:         &fee,
	}

	bodyBytes := marshaler.MustMarshal(body)
	authInfoBytes := marshaler.MustMarshal(authInfo)

	signBytes, err := auxTx.DirectSignBytes(
		bodyBytes,
		authInfoBytes,
		constants.AppChainID,
		0,
	)
	suite.Require().NoError(err)

	return signBytes
}
