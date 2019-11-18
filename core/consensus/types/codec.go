package types

import (
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterBlockAmino(cdc)
}
