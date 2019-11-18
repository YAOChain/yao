package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbm "github.com/tendermint/tm-db"

	abci "github.com/YAOChain/yao/core/abci/types"
	"github.com/YAOChain/yao/core/libs/cli"
	"github.com/YAOChain/yao/core/libs/log"
	tmtypes "github.com/YAOChain/yao/core/types"
	"github.com/YAOChain/yao/node"

	"github.com/YAOChain/yao/app"

	"github.com/YAOChain/yao/baseapp"
	"github.com/YAOChain/yao/client"
	"github.com/YAOChain/yao/store"
	yaotypes "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/x/genaccounts"
	genaccscli "github.com/YAOChain/yao/x/genaccounts/client/cli"
	genutilcli "github.com/YAOChain/yao/x/genutil/client/cli"
	"github.com/YAOChain/yao/x/staking"
)

// gaiad custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := yaotypes.GetConfig()
	config.SetBech32PrefixForAccount(yaotypes.Bech32PrefixAccAddr, yaotypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(yaotypes.Bech32PrefixValAddr, yaotypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(yaotypes.Bech32PrefixConsAddr, yaotypes.Bech32PrefixConsPub)
	config.Seal()

	ctx := node.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "yao",
		Short:             "Yao Network node",
		PersistentPreRunE: node.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx, cdc, app.ModuleBasics, genaccounts.AppModuleBasic{}))
	rootCmd.AddCommand(replayCmd())

	node.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "YAO", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewYaoApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(node.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(node.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		gApp := app.NewYaoApp(logger, db, traceStore, false, uint(1))
		err := gApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	gApp := app.NewYaoApp(logger, db, traceStore, true, uint(1))
	return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
