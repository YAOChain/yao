package consensus

import (
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterConsensusMessages(cdc)
	RegisterWALMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
