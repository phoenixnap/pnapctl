package ranchermodels

import (
	"testing"

	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
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

func assertEqualSshConfig(test_framework *testing.T, cliSshConfig SshConfig, sdkSshConfig ranchersdk.NodePoolSshConfig) {
	assert.Equal(test_framework, cliSshConfig.InstallDefaultKeys, sdkSshConfig.InstallDefaultKeys)

	if testutil.AssertNilEquality(test_framework, "Keys", cliSshConfig.Keys, sdkSshConfig.Keys) {
		assert.Equal(test_framework, len(cliSshConfig.Keys), len(sdkSshConfig.Keys))
		for i := range cliSshConfig.Keys {
			assert.Equal(test_framework, (cliSshConfig.Keys)[i], (sdkSshConfig.Keys)[i])
		}
	}
	if testutil.AssertNilEquality(test_framework, "Key Ids", cliSshConfig.KeyIds, sdkSshConfig.KeyIds) {
		assert.Equal(test_framework, len(cliSshConfig.KeyIds), len(sdkSshConfig.KeyIds))
		for i := range cliSshConfig.KeyIds {
			assert.Equal(test_framework, (cliSshConfig.KeyIds)[i], (sdkSshConfig.KeyIds)[i])
		}
	}
}
