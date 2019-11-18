package types

import (
	"github.com/YAOChain/yao/codec"
)

// module codec
var ModuleCdc = codec.New()

// RegisterCodec registers all the necessary types and interfaces for
// governance.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*Content)(nil), nil)

	cdc.RegisterConcrete(MsgSubmitProposal{}, "yao/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "yao/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgVote{}, "yao/MsgVote", nil)

	cdc.RegisterConcrete(TextProposal{}, "yao/TextProposal", nil)
	cdc.RegisterConcrete(SoftwareUpgradeProposal{}, "yao/SoftwareUpgradeProposal", nil)
}

// RegisterProposalTypeCodec registers an external proposal content type defined
// in another module for the internal ModuleCdc. This allows the MsgSubmitProposal
// to be correctly Amino encoded and decoded.
func RegisterProposalTypeCodec(o interface{}, name string) {
	ModuleCdc.RegisterConcrete(o, name, nil)
}

// TODO determine a good place to seal this codec
func init() {
	RegisterCodec(ModuleCdc)
}
