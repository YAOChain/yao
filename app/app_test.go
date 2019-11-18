package app

import (
	"os"
	"testing"

	"github.com/YAOChain/yao/core/libs/log"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tm-db"

	"github.com/YAOChain/yao/codec"
	"github.com/YAOChain/yao/simapp"

	abci "github.com/YAOChain/yao/core/abci/types"
)

func TestGaiadExport(t *testing.T) {
	db := db.NewMemDB()
	gapp := NewYaoApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	setGenesis(gapp)

	// Making a new app object with the db, so that initchain hasn't been called
	newGapp := NewYaoApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	_, _, err := newGapp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := db.NewMemDB()
	app := NewYaoApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)

	for acc := range maccPerms {
		require.True(t, app.bankKeeper.BlacklistedAddr(app.supplyKeeper.GetModuleAddress(acc)))
	}
}

func setGenesis(gapp *YaoApp) error {

	genesisState := simapp.NewDefaultGenesisState()
	stateBytes, err := codec.MarshalJSONIndent(gapp.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	gapp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	gapp.Commit()
	return nil
}
