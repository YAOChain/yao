package types

import (
	"github.com/YAOChain/yao/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgWithdrawDelegatorReward{}, "yao/MsgWithdrawDelegationReward", nil)
	cdc.RegisterConcrete(MsgWithdrawValidatorCommission{}, "yao/MsgWithdrawValidatorCommission", nil)
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "yao/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(CommunityPoolSpendProposal{}, "yao/CommunityPoolSpendProposal", nil)
}

// generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
