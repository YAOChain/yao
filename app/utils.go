//nolint
package app

import (
	"io"

	"github.com/YAOChain/yao/core/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/YAOChain/yao/baseapp"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/x/staking"
)

var (
	genesisFile        string
	paramsFile         string
	exportParamsPath   string
	exportParamsHeight int
	exportStatePath    string
	exportStatsPath    string
	seed               int64
	initialBlockHeight int
	numBlocks          int
	blockSize          int
	enabled            bool
	verbose            bool
	lean               bool
	commit             bool
	period             int
	onOperation        bool // TODO Remove in favor of binary search for invariant violation
	allInvariants      bool
	genesisTime        int64
)

// DONTCOVER

// NewGaiaAppUNSAFE is used for debugging purposes only.
//
// NOTE: to not use this function with non-test code
func NewGaiaAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*baseapp.BaseApp),
) (gapp *YaoApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	gapp = NewYaoApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return gapp, gapp.keys[baseapp.MainStoreKey], gapp.keys[staking.StoreKey], gapp.stakingKeeper
}
