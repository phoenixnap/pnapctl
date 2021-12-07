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

func assertEqualSshConfig(test_framework *testing.T, s1 SshConfig, s2 ranchersdk.SshConfig) {
	assert.Equal(test_framework, s1.InstallDefaultKeys, s2.InstallDefaultKeys)

	if !assertNilEquality(test_framework, "Keys", s1.Keys, s2.Keys) {
		assert.Equal(test_framework, len(*s1.Keys), len(*s2.Keys))
		for i := range *s1.Keys {
			assert.Equal(test_framework, (*s1.Keys)[i], (*s2.Keys)[i])
		}
	}
	if !assertNilEquality(test_framework, "Key Ids", s1.KeyIds, s2.KeyIds) {
		assert.Equal(test_framework, len(*s1.KeyIds), len(*s2.KeyIds))
		for i := range *s1.KeyIds {
			assert.Equal(test_framework, (*s1.KeyIds)[i], (*s2.KeyIds)[i])
		}
	}
}
