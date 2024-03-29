package cli

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	abciServer "github.com/YAOChain/yao/core/abci/server"
	tcmd "github.com/YAOChain/yao/core/cmd/tendermint/commands"
	"github.com/YAOChain/yao/core/libs/cli"
	"github.com/YAOChain/yao/core/libs/log"
	"github.com/YAOChain/yao/node"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/YAOChain/yao/client"
	"github.com/YAOChain/yao/codec"
	"github.com/YAOChain/yao/node/mock"
	"github.com/YAOChain/yao/tests"
	sdk "github.com/YAOChain/yao/types"
	"github.com/YAOChain/yao/types/module"
	"github.com/YAOChain/yao/x/genutil"
)

var testMbm = module.NewBasicManager(genutil.AppModuleBasic{})

func TestInitCmd(t *testing.T) {
	defer node.SetupViper(t)()
	defer setupClientHome(t)()
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()

	logger := log.NewNopLogger()
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)

	ctx := node.NewContext(cfg, logger)
	cdc := makeCodec()
	cmd := InitCmd(ctx, cdc, testMbm, home)

	require.NoError(t, cmd.RunE(nil, []string{"appnode-test"}))
}

func setupClientHome(t *testing.T) func() {
	clientDir, cleanup := tests.NewTestCaseDir(t)
	viper.Set(flagClientHome, clientDir)
	return cleanup
}

func TestEmptyState(t *testing.T) {
	defer node.SetupViper(t)()
	defer setupClientHome(t)()

	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()

	logger := log.NewNopLogger()
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)

	ctx := node.NewContext(cfg, logger)
	cdc := makeCodec()

	cmd := InitCmd(ctx, cdc, testMbm, home)
	require.NoError(t, cmd.RunE(nil, []string{"appnode-test"}))

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	cmd = node.ExportCmd(ctx, cdc, nil)

	err = cmd.RunE(nil, nil)
	require.NoError(t, err)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	require.Contains(t, out, "genesis_time")
	require.Contains(t, out, "chain_id")
	require.Contains(t, out, "consensus_params")
	require.Contains(t, out, "app_hash")
	require.Contains(t, out, "app_state")
}

func TestStartStandAlone(t *testing.T) {
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()
	viper.Set(cli.HomeFlag, home)
	defer setupClientHome(t)()

	logger := log.NewNopLogger()
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)
	ctx := node.NewContext(cfg, logger)
	cdc := makeCodec()
	initCmd := InitCmd(ctx, cdc, testMbm, home)
	require.NoError(t, initCmd.RunE(nil, []string{"appnode-test"}))

	app, err := mock.NewApp(home, logger)
	require.Nil(t, err)
	svrAddr, _, err := node.FreeTCPAddr()
	require.Nil(t, err)
	svr, err := abciServer.NewServer(svrAddr, "socket", app)
	require.Nil(t, err, "error creating listener")
	svr.SetLogger(logger.With("module", "abci-server"))
	svr.Start()

	timer := time.NewTimer(time.Duration(2) * time.Second)
	select {
	case <-timer.C:
		svr.Stop()
	}
}

func TestInitNodeValidatorFiles(t *testing.T) {
	home, cleanup := tests.NewTestCaseDir(t)
	defer cleanup()
	viper.Set(cli.HomeFlag, home)
	viper.Set(client.FlagName, "moniker")
	cfg, err := tcmd.ParseConfig()
	require.Nil(t, err)
	nodeID, valPubKey, err := genutil.InitializeNodeValidatorFiles(cfg)
	require.Nil(t, err)
	require.NotEqual(t, "", nodeID)
	require.NotEqual(t, 0, len(valPubKey.Bytes()))
}

// custom tx codec
func makeCodec() *codec.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}
