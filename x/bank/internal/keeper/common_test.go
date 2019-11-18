package keeper

// DONTCOVER

import (
	abci "github.com/YAOChain/yao/core/abci/types"
	"github.com/YAOChain/yao/core/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/YAOChain/yao/codec"
	"github.com/YAOChain/yao/store"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/x/auth"
	"github.com/YAOChain/yao/x/bank/internal/types"
	"github.com/YAOChain/yao/x/params"
)

type testInput struct {
	cdc *codec.Codec
	ctx sdk.Context
	k   Keeper
	ak  auth.AccountKeeper
	pk  params.Keeper
}

func setupTestInput() testInput {
	db := dbm.NewMemDB()

	cdc := codec.New()
	auth.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	authCapKey := sdk.NewKVStoreKey("authCapKey")
	keyParams := sdk.NewKVStoreKey("params")
	tkeyParams := sdk.NewTransientStoreKey("transient_params")

	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authCapKey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	ms.LoadLatestVersion()

	blacklistedAddrs := make(map[string]bool)
	blacklistedAddrs[sdk.AccAddress([]byte("moduleAcc")).String()] = true

	pk := params.NewKeeper(cdc, keyParams, tkeyParams, params.DefaultCodespace)

	ak := auth.NewAccountKeeper(
		cdc, authCapKey, pk.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount,
	)
	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain-id"}, false, log.NewNopLogger())

	ak.SetParams(ctx, auth.DefaultParams())

	bankKeeper := NewBaseKeeper(ak, pk.Subspace(types.DefaultParamspace), types.DefaultCodespace, blacklistedAddrs)
	bankKeeper.SetSendEnabled(ctx, true)

	return testInput{cdc: cdc, ctx: ctx, k: bankKeeper, ak: ak, pk: pk}
}
