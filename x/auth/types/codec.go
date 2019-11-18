package types

import (
	"github.com/YAOChain/yao/codec"
	"github.com/YAOChain/yao/x/auth/exported"
)

// RegisterCodec registers concrete types on the codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&BaseAccount{}, "yao/Account", nil)
	cdc.RegisterInterface((*exported.VestingAccount)(nil), nil)
	cdc.RegisterConcrete(&BaseVestingAccount{}, "yao/BaseVestingAccount", nil)
	cdc.RegisterConcrete(&ContinuousVestingAccount{}, "yao/ContinuousVestingAccount", nil)
	cdc.RegisterConcrete(&DelayedVestingAccount{}, "yao/DelayedVestingAccount", nil)
	cdc.RegisterConcrete(StdTx{}, "yao/StdTx", nil)
}

// module wide codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
