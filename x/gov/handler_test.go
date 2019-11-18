package gov

import (
	"strings"
	"testing"

	abci "github.com/YAOChain/yao/core/abci/types"
	sdk "github.com/YAOChain/yao/types"

	"github.com/stretchr/testify/require"
)

func TestInvalidMsg(t *testing.T) {
	k := Keeper{}
	h := NewHandler(k)

	res := h(sdk.NewContext(nil, abci.Header{}, false, nil), sdk.NewTestMsg())
	require.False(t, res.IsOK())
	require.True(t, strings.Contains(res.Log, "unrecognized gov message type"))
}
