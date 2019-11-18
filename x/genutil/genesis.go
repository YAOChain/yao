package genutil

import (
	abci "github.com/YAOChain/yao/core/abci/types"

	"github.com/YAOChain/yao/codec"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(ctx sdk.Context, cdc *codec.Codec, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState GenesisState) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, stakingKeeper, deliverTx)
	}
	return validators
}
