package keys

import (
	"github.com/YAOChain/yao/codec"
)

var cdc *codec.Codec

func init() {
	cdc = codec.New()
	codec.RegisterCrypto(cdc)
	cdc.Seal()
}

// marshal keys
func MarshalJSON(o interface{}) ([]byte, error) {
	return cdc.MarshalJSON(o)
}

// unmarshal json
func UnmarshalJSON(bz []byte, ptr interface{}) error {
	return cdc.UnmarshalJSON(bz, ptr)
}
