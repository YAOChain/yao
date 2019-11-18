package evidence

import (
	cryptoAmino "github.com/YAOChain/yao/core/crypto/encoding/amino"
	"github.com/YAOChain/yao/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	RegisterEvidenceMessages(cdc)
	cryptoAmino.RegisterAmino(cdc)
	types.RegisterEvidences(cdc)
}

// For testing purposes only
func RegisterMockEvidences() {
	types.RegisterMockEvidences(cdc)
}
