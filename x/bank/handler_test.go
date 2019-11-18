package bank

import (
	"strings"
	"testing"

	abci "github.com/YAOChain/yao/core/abci/types"
	sdk "github.com/YAOChain/yao/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	h := NewHandler(nil)

	res := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized bank message type"))
}
