package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
	"github.com/stretchr/testify/assert"
)

func TestSshConfigToSdk(test_framework *testing.T) {
	sshConfig := GenerateSshConfigCli()
	sdkSshConfig := *sshConfig.ToSdk()

	assertEqualSshConfig(test_framework, sshConfig, sdkSshConfig)
}

func TestSshConfigFromSdk(test_framework *testing.T) {
	sdkSshConfig := GenerateSshConfigSdk()
	sshConfig := *SshConfigFromSdk(&sdkSshConfig)

	assertEqualSshConfig(test_framework, sshConfig, sdkSshConfig)
}

func assertEqualSshConfig(test_framework *testing.T, cliSshConfig SshConfig, sdkSshConfig ranchersdk.SshConfig) {
	assert.Equal(test_framework, cliSshConfig.InstallDefaultKeys, sdkSshConfig.InstallDefaultKeys)

	if !assertNilEquality(test_framework, "Keys", cliSshConfig.Keys, sdkSshConfig.Keys) {
		assert.Equal(test_framework, len(*cliSshConfig.Keys), len(*sdkSshConfig.Keys))
		for i := range *cliSshConfig.Keys {
			assert.Equal(test_framework, (*cliSshConfig.Keys)[i], (*sdkSshConfig.Keys)[i])
		}
	}
	if !assertNilEquality(test_framework, "Key Ids", cliSshConfig.KeyIds, sdkSshConfig.KeyIds) {
		assert.Equal(test_framework, len(*cliSshConfig.KeyIds), len(*sdkSshConfig.KeyIds))
		for i := range *cliSshConfig.KeyIds {
			assert.Equal(test_framework, (*cliSshConfig.KeyIds)[i], (*sdkSshConfig.KeyIds)[i])
		}
	}
}
