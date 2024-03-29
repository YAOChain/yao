package rest

import (
	"net/http"

	"github.com/YAOChain/yao/client/context"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/types/rest"
	"github.com/YAOChain/yao/x/genutil/types"
)

// QueryGenesisTxs writes the genesis transactions to the response if no error
// occurs.
func QueryGenesisTxs(cliCtx context.CLIContext, w http.ResponseWriter) {
	resultGenesis, err := cliCtx.Client.Genesis()
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError,
			sdk.AppendMsgToErr("could not retrieve genesis from client", err.Error()))
		return
	}

	appState, err := types.GenesisStateFromGenDoc(cliCtx.Codec, *resultGenesis.Genesis)
	if err != nil {
		rest.WriteErrorResponse(w, http.StatusInternalServerError,
			sdk.AppendMsgToErr("could not decode genesis doc", err.Error()))
		return
	}

	genState := types.GetGenesisStateFromAppState(cliCtx.Codec, appState)
	genTxs := make([]sdk.Tx, len(genState.GenTxs))
	for i, tx := range genState.GenTxs {
		err := cliCtx.Codec.UnmarshalJSON(tx, &genTxs[i])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError,
				sdk.AppendMsgToErr("could not decode genesis transaction", err.Error()))
			return
		}
	}

	rest.PostProcessResponseBare(w, cliCtx, genTxs)
}
