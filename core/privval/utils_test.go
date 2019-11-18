package privval

import (
	"fmt"
	"testing"

	cmn "github.com/YAOChain/yao/core/libs/common"
	"github.com/stretchr/testify/assert"
)

func TestIsConnTimeoutForNonTimeoutErrors(t *testing.T) {
	assert.False(t, IsConnTimeout(cmn.ErrorWrap(ErrDialRetryMax, "max retries exceeded")))
	assert.False(t, IsConnTimeout(fmt.Errorf("completely irrelevant error")))
}
