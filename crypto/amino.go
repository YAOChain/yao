package crypto

import (
	cryptoAmino "github.com/YAOChain/yao/core/crypto/encoding/amino"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterAmino(cdc)
	cryptoAmino.RegisterAmino(cdc)
}

// RegisterAmino registers all go-crypto related types in the given (amino) codec.
func RegisterAmino(cdc *amino.Codec) {
	cdc.RegisterConcrete(PrivKeyLedgerSecp256k1{},
		"yao/PrivKeyLedgerSecp256k1", nil)
}
