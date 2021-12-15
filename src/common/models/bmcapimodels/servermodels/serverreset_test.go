package servermodels

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestServerResestToSdk(test_framework *testing.T) {

	cliModel := GenerateServerResetCli()
	sdkModel := ServerResetToSDK(&cliModel)

	assert.Equal(test_framework, sdkModel.InstallDefaultSshKeys, cliModel.InstallDefaultSshKeys)
	assert.Equal(test_framework, sdkModel.SshKeys, cliModel.SshKeys)
	assert.Equal(test_framework, sdkModel.SshKeyIds, cliModel.SshKeyIds)

	if testutil.AssertNilEquality(test_framework, "OsConfiguration", sdkModel.OsConfiguration, cliModel.OsConfiguration) {
		assertEqualOsConfigurationMap(test_framework, *cliModel.OsConfiguration, *sdkModel.OsConfiguration)
	}
}

func TestNilServerResestToSdk(test_framework *testing.T) {

	var cliModel *ServerReset = nil
	sdkModel := ServerResetToSDK(cliModel)

	assert.Nil(test_framework, sdkModel)
}
