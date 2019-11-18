package types

import (
	"github.com/YAOChain/yao/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "yao/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "yao/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "yao/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "yao/MsgUndelegate", nil)
	cdc.RegisterConcrete(MsgBeginRedelegate{}, "yao/MsgBeginRedelegate", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
