package types

import (
	"github.com/YAOChain/yao/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "yao/MsgSend", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "yao/MsgMultiSend", nil)
}

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	ModuleCdc.Seal()
}
