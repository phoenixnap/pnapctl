package ipmodels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpBlockFromSdk(test_framework *testing.T) {
	sdkModel := GenerateIpBlockSdk()
	cliModel := IpBlockFromSdk(sdkModel)

	assert.Equal(test_framework, sdkModel.Id, cliModel.Id)
	assert.Equal(test_framework, sdkModel.Location, cliModel.Location)
	assert.Equal(test_framework, sdkModel.CidrBlockSize, cliModel.CidrBlockSize)
	assert.Equal(test_framework, sdkModel.Cidr, cliModel.Cidr)
	assert.Equal(test_framework, sdkModel.Status, cliModel.Status)
	assert.Equal(test_framework, sdkModel.AssignedResourceId, cliModel.AssignedResourceId)
	assert.Equal(test_framework, sdkModel.AssignedResourceType, cliModel.AssignedResourceType)
}
