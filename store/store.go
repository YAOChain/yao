package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/YAOChain/yao/store/rootmulti"
	"github.com/YAOChain/yao/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewPruningOptionsFromString(strategy string) (opt PruningOptions) {
	switch strategy {
	case "nothing":
		opt = PruneNothing
	case "everything":
		opt = PruneEverything
	case "syncable":
		opt = PruneSyncable
	default:
		opt = PruneSyncable
	}
	return
}
