package client

import (
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	types.RegisterEvidences(cdc)
}
