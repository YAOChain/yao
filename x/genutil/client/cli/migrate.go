package cli

import (
	"fmt"
	"time"

	"github.com/YAOChain/yao/core/types"
	"github.com/YAOChain/yao/node"
	"github.com/spf13/cobra"

	"github.com/YAOChain/yao/codec"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/version"
	extypes "github.com/YAOChain/yao/x/genutil"
	v036 "github.com/YAOChain/yao/x/genutil/legacy/v036"
)

var migrationMap = extypes.MigrationMap{
	"v0.36": v036.Migrate,
}

const (
	flagGenesisTime = "genesis-time"
	flagChainId     = "chain-id"
)

func MigrateGenesisCmd(_ *node.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate [target-version] [genesis-file]",
		Short: "Migrate genesis to a specified target version",
		Long: fmt.Sprintf(`Migrate the source genesis into the target version and print to STDOUT.

Example:
$ %s migrate v0.36 /path/to/genesis.json --chain-id=cosmoshub-3 --genesis-time=2019-04-22T17:00:00Z
`, version.ServerName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			target := args[0]
			importGenesis := args[1]

			genDoc, err := types.GenesisDocFromFile(importGenesis)
			if err != nil {
				return err
			}

			var initialState extypes.AppMap
			cdc.MustUnmarshalJSON(genDoc.AppState, &initialState)

			if migrationMap[target] == nil {
				return fmt.Errorf("unknown migration function version: %s", target)
			}

			newGenState := migrationMap[target](initialState)
			genDoc.AppState = cdc.MustMarshalJSON(newGenState)

			genesisTime := cmd.Flag(flagGenesisTime).Value.String()
			if genesisTime != "" {
				var t time.Time

				err := t.UnmarshalText([]byte(genesisTime))
				if err != nil {
					return err
				}

				genDoc.GenesisTime = t
			}

			chainId := cmd.Flag(flagChainId).Value.String()
			if chainId != "" {
				genDoc.ChainID = chainId
			}

			out, err := cdc.MarshalJSONIndent(genDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sdk.MustSortJSON(out)))
			return nil
		},
	}

	cmd.Flags().String(flagGenesisTime, "", "Override genesis_time with this flag")
	cmd.Flags().String(flagChainId, "", "Override chain_id with this flag")

	return cmd
}
