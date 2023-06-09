package gosdk

import (
	"fmt"
	"github.com/okx/okbchain-go-sdk/module/feesplit"
	"github.com/okx/okbchain-go-sdk/module/ibc"
	"github.com/okx/okbchain-go-sdk/module/wasm"
	ibcTypes "github.com/okx/okbchain/libs/ibc-go/modules/apps/transfer/types"
	feesplitTypes "github.com/okx/okbchain/x/feesplit/types"

	"github.com/okx/okbchain-go-sdk/exposed"
	"github.com/okx/okbchain-go-sdk/module"
	"github.com/okx/okbchain-go-sdk/module/auth"
	authtypes "github.com/okx/okbchain-go-sdk/module/auth/types"
	"github.com/okx/okbchain-go-sdk/module/distribution"
	distrtypes "github.com/okx/okbchain-go-sdk/module/distribution/types"
	"github.com/okx/okbchain-go-sdk/module/evm"
	evmtypes "github.com/okx/okbchain-go-sdk/module/evm/types"
	"github.com/okx/okbchain-go-sdk/module/governance"
	govtypes "github.com/okx/okbchain-go-sdk/module/governance/types"
	"github.com/okx/okbchain-go-sdk/module/slashing"
	slashingtypes "github.com/okx/okbchain-go-sdk/module/slashing/types"
	"github.com/okx/okbchain-go-sdk/module/staking"
	stakingtypes "github.com/okx/okbchain-go-sdk/module/staking/types"
	"github.com/okx/okbchain-go-sdk/module/tendermint"
	tmtypes "github.com/okx/okbchain-go-sdk/module/tendermint/types"
	"github.com/okx/okbchain-go-sdk/module/token"
	tokentypes "github.com/okx/okbchain-go-sdk/module/token/types"
	gosdktypes "github.com/okx/okbchain-go-sdk/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/codec"
	wasmTypes "github.com/okx/okbchain/x/wasm/types"
)

// Client - structure of the main client of OKBChain GoSDK
type Client struct {
	config  gosdktypes.ClientConfig
	cdc     *codec.Codec
	modules map[string]gosdktypes.Module
}

// NewClient creates a new instance of Client
func NewClient(config gosdktypes.ClientConfig) Client {
	cdc := gosdktypes.NewCodec()
	pClient := &Client{
		config:  config,
		cdc:     cdc,
		modules: make(map[string]gosdktypes.Module),
	}
	pBaseClient := module.NewBaseClient(cdc, &pClient.config)

	pClient.registerModule(
		auth.NewAuthClient(pBaseClient),
		distribution.NewDistrClient(pBaseClient),
		evm.NewEvmClient(pBaseClient),
		governance.NewGovClient(pBaseClient),
		staking.NewStakingClient(pBaseClient),
		slashing.NewSlashingClient(pBaseClient),
		token.NewTokenClient(pBaseClient),
		tendermint.NewTendermintClient(pBaseClient),
		ibc.NewIbcClient(pBaseClient),
		feesplit.NewfeesplitClient(pBaseClient),
		wasm.NewWasmClient(pBaseClient),
	)

	return *pClient
}

func (cli *Client) registerModule(mods ...gosdktypes.Module) {
	for _, mod := range mods {
		moduleName := mod.Name()
		if _, ok := cli.modules[moduleName]; ok {
			panic(fmt.Sprintf("duplicated module: %s", moduleName))
		}
		// register codec by each module
		mod.RegisterCodec(cli.cdc)
		cli.modules[moduleName] = mod
	}
	gosdktypes.RegisterBasicCodec(cli.cdc)
	cli.cdc.Seal()
}

// GetConfig returns the client config
func (cli *Client) GetConfig() gosdktypes.ClientConfig {
	return cli.config
}

func (cli *Client) Auth() exposed.Auth {
	return cli.modules[authtypes.ModuleName].(exposed.Auth)
}

func (cli *Client) Distribution() exposed.Distribution {
	return cli.modules[distrtypes.ModuleName].(exposed.Distribution)
}
func (cli *Client) Evm() exposed.Evm {
	return cli.modules[evmtypes.ModuleName].(exposed.Evm)
}

func (cli *Client) Governance() exposed.Governance {
	return cli.modules[govtypes.ModuleName].(exposed.Governance)
}

func (cli *Client) Slashing() exposed.Slashing {
	return cli.modules[slashingtypes.ModuleName].(exposed.Slashing)
}
func (cli *Client) Staking() exposed.Staking {
	return cli.modules[stakingtypes.ModuleName].(exposed.Staking)
}
func (cli *Client) Tendermint() exposed.Tendermint {
	return cli.modules[tmtypes.ModuleName].(exposed.Tendermint)
}
func (cli *Client) Token() exposed.Token {
	return cli.modules[tokentypes.ModuleName].(exposed.Token)
}

func (cli *Client) Ibc() exposed.Ibc {
	return cli.modules[ibcTypes.ModuleName].(exposed.Ibc)
}

func (cli *Client) Feesplit() exposed.Feesplit {
	return cli.modules[feesplitTypes.ModuleName].(exposed.Feesplit)
}

func (cli *Client) Wasm() exposed.Wasm {
	return cli.modules[wasmTypes.ModuleName].(exposed.Wasm)
}
