package privval

import (
	"testing"
	"time"

	"github.com/YAOChain/yao/core/crypto/ed25519"
	cmn "github.com/YAOChain/yao/core/libs/common"
	"github.com/stretchr/testify/assert"
)

func TestIsConnTimeoutForFundamentalTimeouts(t *testing.T) {
	// Generate a networking timeout
	dialer := DialTCPFn(testFreeTCPAddr(t), time.Millisecond, ed25519.GenPrivKey())
	_, err := dialer()
	assert.Error(t, err)
	assert.True(t, IsConnTimeout(err))
}

func TestIsConnTimeoutForWrappedConnTimeouts(t *testing.T) {
	dialer := DialTCPFn(testFreeTCPAddr(t), time.Millisecond, ed25519.GenPrivKey())
	_, err := dialer()
	assert.Error(t, err)
	err = cmn.ErrorWrap(ErrConnTimeout, err.Error())
	assert.True(t, IsConnTimeout(err))
}
