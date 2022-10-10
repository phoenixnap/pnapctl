package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestOsConfigurationMapEsxiToSdk(test_framework *testing.T) {
	cliModel := GenerateOsConfigurationMapEsxiCli()
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationMapEsxi(test_framework, *cliModel, *sdkModel)
}

func TestEmptyOsConfigurationMapEsxiToSdk(test_framework *testing.T) {
	cliModel := &OsConfigurationMapEsxi{}
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationMapEsxi(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationMapEsxiToSdk(test_framework *testing.T) {
	var cliModel *OsConfigurationMapEsxi = nil

	assert.Nil(test_framework, cliModel.toSdk())
}

// assertion functions
func assertEqualOsConfigurationMapEsxi(test_framework *testing.T, cliOsConfigurationMapEsxi OsConfigurationMapEsxi, sdkOsConfigurationMapEsxi bmcapisdk.OsConfigurationMapEsxi) {

	assert.Equal(test_framework, cliOsConfigurationMapEsxi.RootPassword, sdkOsConfigurationMapEsxi.RootPassword)
	assert.Equal(test_framework, cliOsConfigurationMapEsxi.ManagementUiUrl, sdkOsConfigurationMapEsxi.ManagementUiUrl)

	if testutil.AssertNilEquality(test_framework, "Management Access Allowed IPs", cliOsConfigurationMapEsxi.ManagementAccessAllowedIps, sdkOsConfigurationMapEsxi.ManagementAccessAllowedIps) {
		assert.Equal(test_framework, len(cliOsConfigurationMapEsxi.ManagementAccessAllowedIps), len(sdkOsConfigurationMapEsxi.ManagementAccessAllowedIps))

		for i := range cliOsConfigurationMapEsxi.ManagementAccessAllowedIps {
			assert.Equal(test_framework, (cliOsConfigurationMapEsxi.ManagementAccessAllowedIps)[i], (sdkOsConfigurationMapEsxi.ManagementAccessAllowedIps)[i])
		}
	}
}
