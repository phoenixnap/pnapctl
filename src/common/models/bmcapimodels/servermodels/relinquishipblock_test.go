package servermodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelinquishIpBlockToSdk(test_framework *testing.T) {

	cliModel := GenerateRelinquishIpBlockCli()
	sdkModel := cliModel.ToSdk()

	assert.Equal(test_framework, cliModel.DeleteIpBlocks, sdkModel.DeleteIpBlocks)
}
