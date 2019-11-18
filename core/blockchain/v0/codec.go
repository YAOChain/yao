package v0

import (
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
