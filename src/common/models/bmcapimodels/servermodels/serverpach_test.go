package servermodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerPatchToSDK(test_framework *testing.T) {
	cliModel := GenerateServerPatchCli()
	sdkModel := cliModel.ToSdk()

	assert.Equal(test_framework, sdkModel.Hostname, cliModel.Hostname)
	assert.Equal(test_framework, sdkModel.Description, cliModel.Description)
}
