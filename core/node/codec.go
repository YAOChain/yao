package node

import (
	cryptoAmino "github.com/YAOChain/yao/core/crypto/encoding/amino"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}
