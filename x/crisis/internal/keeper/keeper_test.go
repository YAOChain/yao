package keeper

import (
	"testing"

	"github.com/YAOChain/yao/core/libs/log"
	"github.com/stretchr/testify/require"

	"github.com/YAOChain/yao/codec"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/x/crisis/internal/types"
	"github.com/YAOChain/yao/x/params"
)

func testPassingInvariant(_ sdk.Context) (string, bool) {
	return "", false
}

func testFailingInvariant(_ sdk.Context) (string, bool) {
	return "", true
}

func testKeeper(checkPeriod uint) Keeper {
	cdc := codec.New()
	paramsKeeper := params.NewKeeper(
		cdc, sdk.NewKVStoreKey(params.StoreKey), sdk.NewTransientStoreKey(params.TStoreKey), params.DefaultCodespace,
	)

	return NewKeeper(paramsKeeper.Subspace(types.DefaultParamspace), checkPeriod, nil, "test")
}

func TestLogger(t *testing.T) {
	k := testKeeper(5)

	ctx := sdk.Context{}.WithLogger(log.NewNopLogger())
	require.Equal(t, ctx.Logger(), k.Logger(ctx))
}

func TestInvariants(t *testing.T) {
	k := testKeeper(5)
	require.Equal(t, k.InvCheckPeriod(), uint(5))

	k.RegisterRoute("testModule", "testRoute", testPassingInvariant)
	require.Len(t, k.Routes(), 1)
}

func TestAssertInvariants(t *testing.T) {
	k := testKeeper(5)
	ctx := sdk.Context{}.WithLogger(log.NewNopLogger())

	k.RegisterRoute("testModule", "testRoute1", testPassingInvariant)
	require.NotPanics(t, func() { k.AssertInvariants(ctx) })

	k.RegisterRoute("testModule", "testRoute2", testFailingInvariant)
	require.Panics(t, func() { k.AssertInvariants(ctx) })
}
