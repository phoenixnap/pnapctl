package servermodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerCreateToSDK(test_framework *testing.T) {
	cliModel := GenerateServerCreateCli()
	sdkModel := cliModel.ToSdk()

	assert.Equal(test_framework, cliModel.Hostname, sdkModel.Hostname)
	assert.Equal(test_framework, cliModel.Description, sdkModel.Description)
	assert.Equal(test_framework, cliModel.Os, sdkModel.Os)
	assert.Equal(test_framework, cliModel.Location, sdkModel.Location)
	assert.Equal(test_framework, cliModel.InstallDefaultSshKeys, sdkModel.InstallDefaultSshKeys)
	assert.Equal(test_framework, cliModel.SshKeys, sdkModel.SshKeys)
	assert.Equal(test_framework, cliModel.SshKeyIds, sdkModel.SshKeyIds)
	assert.Equal(test_framework, cliModel.ReservationId, sdkModel.ReservationId)
	assert.Equal(test_framework, cliModel.PricingModel, sdkModel.PricingModel)
	assert.Equal(test_framework, cliModel.NetworkType, sdkModel.NetworkType)

	assertEqualOsConfiguration(test_framework, *cliModel.OsConfiguration, *sdkModel.OsConfiguration)

	sdkTags := *sdkModel.Tags
	for i, tagAssignmentRequest := range *cliModel.Tags {
		assertEqualTagAssignmentRequest(test_framework, tagAssignmentRequest, sdkTags[i])
	}

	assertEqualNetworkConfiguration(test_framework, *cliModel.NetworkConfiguration, *sdkModel.NetworkConfiguration)

}
