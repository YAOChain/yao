package slashing

import (
	"testing"

	abci "github.com/YAOChain/yao/core/abci/types"
	"github.com/stretchr/testify/require"

	"github.com/YAOChain/yao/codec"
)

func TestNewQuerier(t *testing.T) {
	ctx, _, _, _, keeper := createTestInput(t, keeperTestParams())
	querier := NewQuerier(keeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(ctx, []string{"parameters"}, query)
	require.NoError(t, err)
}

func TestQueryParams(t *testing.T) {
	cdc := codec.New()
	ctx, _, _, _, keeper := createTestInput(t, keeperTestParams())

	var params Params

	res, errRes := queryParams(ctx, keeper)
	require.NoError(t, errRes)

	err := cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)
	require.Equal(t, keeper.GetParams(ctx), params)
}
