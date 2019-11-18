package main

import (
	ctypes "github.com/YAOChain/yao/core/rpc/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
