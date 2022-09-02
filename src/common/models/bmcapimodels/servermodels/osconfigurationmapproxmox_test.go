package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// tests
func TestOsConfigurationMapProxmoxToSdk(test_framework *testing.T) {
	cliModel := GenerateOsConfigurationMapProxmoxCli()
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationMapProxmox(test_framework, *cliModel, *sdkModel)
}

func TestEmptyOsConfigurationMapProxmoxToSdk(test_framework *testing.T) {
	cliModel := &OsConfigurationMapProxmox{}
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationMapProxmox(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationMapProxmoxToSdk(test_framework *testing.T) {
	var cliModel *OsConfigurationMapProxmox = nil

	assert.Nil(test_framework, cliModel.toSdk())
}

// assertion functions
func assertEqualOsConfigurationMapProxmox(test_framework *testing.T, cliOsConfigurationMapProxmox OsConfigurationMapProxmox, sdkOsConfigurationMapProxmox bmcapisdk.OsConfigurationMapProxmox) {

	assert.Equal(test_framework, cliOsConfigurationMapProxmox.RootPassword, sdkOsConfigurationMapProxmox.RootPassword)
	assert.Equal(test_framework, cliOsConfigurationMapProxmox.ManagementUiUrl, sdkOsConfigurationMapProxmox.ManagementUiUrl)

	if testutil.AssertNilEquality(test_framework, "Management Access Allowed IPs", cliOsConfigurationMapProxmox.ManagementAccessAllowedIps, sdkOsConfigurationMapProxmox.ManagementAccessAllowedIps) {
		assert.Equal(test_framework, len(cliOsConfigurationMapProxmox.ManagementAccessAllowedIps), len(sdkOsConfigurationMapProxmox.ManagementAccessAllowedIps))

		for i := range cliOsConfigurationMapProxmox.ManagementAccessAllowedIps {
			assert.Equal(test_framework, (cliOsConfigurationMapProxmox.ManagementAccessAllowedIps)[i], (sdkOsConfigurationMapProxmox.ManagementAccessAllowedIps)[i])
		}
	}
}
