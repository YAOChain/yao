package core_types

import (
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

func RegisterAmino(cdc *amino.Codec) {
	types.RegisterEventDatas(cdc)
	types.RegisterBlockAmino(cdc)
}
