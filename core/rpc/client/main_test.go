package client_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/YAOChain/yao/core/abci/example/kvstore"
	nm "github.com/YAOChain/yao/core/node"
	rpctest "github.com/YAOChain/yao/core/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a tendermint node (and kvstore) in the background to test against
	dir, err := ioutil.TempDir("/tmp", "rpc-client-test")
	if err != nil {
		panic(err)
	}
	app := kvstore.NewPersistentKVStoreApplication(dir)
	node = rpctest.StartTendermint(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopTendermint(node)
	os.Exit(code)
}
