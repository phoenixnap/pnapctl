package ipmodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpBlockCreateToSdk(test_framework *testing.T) {
	cliModel := GenerateIpBlockCreateCLI()
	sdkModel := cliModel.ToSdk()

	assert.Equal(test_framework, cliModel.CidrBlockSize, sdkModel.CidrBlockSize)
	assert.Equal(test_framework, cliModel.Location, sdkModel.Location)
}
