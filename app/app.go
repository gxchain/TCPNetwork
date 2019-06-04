package app

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/cosmos/cosmos-sdk/x/bank"

	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/gxchain/TCPNetwork/x/tcp"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	appName = "tcp"
	TCPStoreKey = "tcp"

	// The ModuleBasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
)

type tcpApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	// keys to access the substores
	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyStaking       *sdk.KVStoreKey
	tkeyStaking      *sdk.TransientStoreKey
	keuySlashing	 *sdk.KVStore

	keyDistr         *sdk.KVStoreKey
	tkeyDistr        *sdk.TransientStoreKey

	keyGov			 *sdk.KVStore
	keyTCP           *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey

	// keepers
	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper

	bankKeeper          bank.Keeper
	stakingKeeper       staking.Keeper
	slashingKeeper 		slashing.Keeper

	distrKeeper         distr.Keeper
	govKeeper 			gov.Keeper

	paramsKeeper        params.Keeper
	tcpKeeper           tcp.Keeper

	// the module manager
	mm *sdk.ModuleManager
}

// NewTCPApp is a constructor function for tcpApp
func NewTCPApp(logger log.Logger, db dbm.DB) *tcpApp {

	// First define the top level codec that will be shared by the different modules
	cdc := CreateCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	// Here you initialize your application with the store keys it requires
	var app = &tcpApp{
		BaseApp: bApp,
		cdc:     cdc,

		keyTCP:           sdk.NewKVStoreKey(TCPStoreKey),
		keyMain:          sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccount:       sdk.NewKVStoreKey(auth.StoreKey),
		keyStaking:       sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking:      sdk.NewTransientStoreKey(staking.TStoreKey),
		//keyMint:          sdk.NewKVStoreKey(mint.StoreKey),
		keyDistr:         sdk.NewKVStoreKey(distr.StoreKey),
		//tkeyDistr:        sdk.NewTransientStoreKey(distr.TStoreKey),
		//keySlashing:      sdk.NewKVStoreKey(slashing.StoreKey),
		//keyGov:           sdk.NewKVStoreKey(gov.StoreKey),
		keyFeeCollection: sdk.NewKVStoreKey(auth.FeeStoreKey),
		keyParams:        sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:       sdk.NewTransientStoreKey(params.TStoreKey),
	}

	// init params keeper
	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)

	// add keepers
	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	// The FeeCollectionKeeper collects transaction fees and renders them to the fee distribution module
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	app.stakingKeeper = staking.NewKeeper(app.cdc, app.keyStaking, app.tkeyStaking, app.bankKeeper,
		app.paramsKeeper.Subspace(staking.DefaultParamspace), staking.DefaultCodespace)

	// The TCPKeeper is the Keeper from the module for this tutorial
	// It handles interactions with the tcp
	app.tcpKeeper = tcp.NewKeeper(
		app.bankKeeper,
		app.keyTCP,
		app.cdc,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	//app.stakingKeeper = *stakingKeeper.SetHooks(
	//	staking.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()))



	// The app.Router is the main transaction router where each module registers its routes
	// Register the bank and tcp routes here
	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("tcp", tcp.NewHandler(app.tcpKeeper))

	// The app.QueryRouter is the main query router where each module registers its routes
	app.QueryRouter().
		AddRoute("acc", auth.NewQuerier(app.accountKeeper)).
		AddRoute("tcp", tcp.NewQuerier(app.tcpKeeper))


	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyTCP,
		app.keyFeeCollection,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

func (app *tcpApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(GenesisState)
	err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		panic(err)
	}

	// load the accounts
	for _, acc := range genesisState.Accounts {
		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)

	return abci.ResponseInitChain{}
}

// ExportAppStateAndValidators does the things
func (app *tcpApp) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abci.Header{})
	accounts := []*auth.BaseAccount{}

	appendAccountsFn := func(acc auth.Account) bool {
		account := &auth.BaseAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}

		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := GenesisState{
		Accounts: accounts,
		AuthData: auth.DefaultGenesisState(),
		BankData: bank.DefaultGenesisState(),
	}

	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, err
}

func init() {
	ModuleBasics = sdk.NewModuleBasicManager(
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
	)
}

// CreateCodec generates the necessary codecs for Amino
func CreateCodec() *codec.Codec {
	var cds = codec.New()
	MOduleBasics.Re
	//var cdc = codec.New()
	//auth.RegisterCodec(cdc)
	//bank.RegisterCodec(cdc)
	//tcp.RegisterCodec(cdc)
	//staking.RegisterCodec(cdc)
	//sdk.RegisterCodec(cdc)
	//codec.RegisterCrypto(cdc)
	//return cdc
}

// BeginBlocker signals the beginning of a block. It performs application
// updates on the start of every block.
func (app *tcpApp) BeginBlocker(
	_ sdk.Context, _ abci.RequestBeginBlock,
) abci.ResponseBeginBlock {

	return abci.ResponseBeginBlock{}
}

// EndBlocker signals the end of a block. It performs application updates on
// the end of every block.
func (app *tcpApp) EndBlocker(
	_ sdk.Context, _ abci.RequestEndBlock,
) abci.ResponseEndBlock {

	return abci.ResponseEndBlock{}
}

// load a particular height
func (app *tcpApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}
