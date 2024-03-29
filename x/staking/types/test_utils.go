package types

import (
	"github.com/YAOChain/yao/core/crypto"
	"github.com/YAOChain/yao/core/crypto/ed25519"

	sdk "github.com/YAOChain/yao/types"
)

// nolint: deadcode unused
var (
	pk1      = ed25519.GenPrivKey().PubKey()
	pk2      = ed25519.GenPrivKey().PubKey()
	pk3      = ed25519.GenPrivKey().PubKey()
	addr1    = pk1.Address()
	addr2    = pk2.Address()
	addr3    = pk3.Address()
	valAddr1 = sdk.ValAddress(addr1)
	valAddr2 = sdk.ValAddress(addr2)
	valAddr3 = sdk.ValAddress(addr3)

	emptyAddr   sdk.ValAddress
	emptyPubkey crypto.PubKey
)
