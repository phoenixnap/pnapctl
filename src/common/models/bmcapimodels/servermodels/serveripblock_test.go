package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

// tests
func TestServerIpBlockToSdk(test_framework *testing.T) {
	cliModel := GenerateServerIpBlockCli()
	sdkModel := cliModel.ToSdk()

	assert.Equal(test_framework, cliModel.Id, sdkModel.Id)
	assert.Equal(test_framework, &cliModel.VlanId, &sdkModel.VlanId)
}

func TestNilMapServerIpBlocksToSdk(test_framework *testing.T) {
	assert.Nil(test_framework, mapServerIpBlocksToSdk(nil))
}

func TestMapServerIpBlocksToSdk(test_framework *testing.T) {
	cliModels := GenerateServerIpBlockListCli(3)
	sdkModels := mapServerIpBlocksToSdk(cliModels)

	assertServerIpBlockListEquality(test_framework, cliModels, sdkModels)
}

func TestNilMapServerIpBlocksToCLI(test_framework *testing.T) {
	assert.Nil(test_framework, mapServerIpBlocksToCLI(nil))
}

func TestMapServerIpBlocksToCLI(test_framework *testing.T) {
	sdkModels := GenerateServerIpBlockListSdk(3)
	cliModels := mapServerIpBlocksToCLI(sdkModels)
	assertServerIpBlockListEquality(test_framework, cliModels, sdkModels)
}

func assertServerIpBlockListEquality(test_framework *testing.T, cliModels *[]ServerIpBlock, sdkModels *[]bmcapisdk.ServerIpBlock) {
	assert.Equal(test_framework, len(*sdkModels), len(*cliModels))
	for i, cliElement := range *cliModels {
		assert.Equal(test_framework, cliElement.Id, (*sdkModels)[i].Id)
		assert.Equal(test_framework, cliElement.VlanId, (*sdkModels)[i].VlanId)
	}
}
