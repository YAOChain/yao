package cli

import (
	"encoding/json"
	"path/filepath"

	"github.com/YAOChain/yao/core/libs/cli"
	tmtypes "github.com/YAOChain/yao/core/types"
	"github.com/YAOChain/yao/node"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/YAOChain/yao/client"
	"github.com/YAOChain/yao/codec"
	"github.com/YAOChain/yao/x/genutil"
	"github.com/YAOChain/yao/x/genutil/types"
)

const flagGenTxDir = "gentx-dir"

// CollectGenTxsCmd - return the cobra command to collect genesis transactions
func CollectGenTxsCmd(ctx *node.Context, cdc *codec.Codec,
	genAccIterator types.GenesisAccountsIterator, defaultNodeHome string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "collect-gentxs",
		Short: "Collect genesis txs and output a genesis.json file",
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))
			name := viper.GetString(client.FlagName)
			nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return err
			}

			genTxsDir := viper.GetString(flagGenTxDir)
			if genTxsDir == "" {
				genTxsDir = filepath.Join(config.RootDir, "config", "gentx")
			}

			toPrint := newPrintInfo(config.Moniker, genDoc.ChainID, nodeID, genTxsDir, json.RawMessage(""))
			initCfg := genutil.NewInitConfig(genDoc.ChainID, genTxsDir, name, nodeID, valPubKey)

			appMessage, err := genutil.GenAppStateFromConfig(cdc, config, initCfg, *genDoc, genAccIterator)
			if err != nil {
				return err
			}

			toPrint.AppMessage = appMessage

			// print out some key information
			return displayInfo(cdc, toPrint)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagGenTxDir, "",
		"override default \"gentx\" directory from which collect and execute "+
			"genesis transactions; default [--home]/config/gentx/")
	return cmd
}

// DONTCOVER
