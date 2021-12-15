package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestOsConfigurationToSdk(test_framework *testing.T) {
	cliModel := GenerateOsConfigurationCli()
	sdkModel := cliModel.toSdk()

	assertEqualOsConfiguration(test_framework, cliModel, *sdkModel)
}

func TestEmptyOsConfigurationToSdk(test_framework *testing.T) {
	var cliModel *OsConfiguration = &OsConfiguration{}
	sdkModel := cliModel.toSdk()

	assertEqualOsConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationToSdk(test_framework *testing.T) {
	var cliModel *OsConfiguration = nil

	assert.Nil(test_framework, cliModel.toSdk())
}

func TestOsConfigurationFromSdk(test_framework *testing.T) {
	sdkModel := GenerateOsConfigurationSdk()
	cliModel := OsConfigurationFromSdk(&sdkModel)

	assertEqualOsConfiguration(test_framework, *cliModel, sdkModel)
}

func TestEmptyOsConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfiguration = &bmcapisdk.OsConfiguration{}
	cliModel := OsConfigurationFromSdk(sdkModel)

	assertEqualOsConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfiguration = nil

	assert.Nil(test_framework, OsConfigurationFromSdk(sdkModel))
}

// assertion functions
func assertEqualOsConfiguration(test_framework *testing.T, cliOsConfiguration OsConfiguration, sdkOsConfiguration bmcapisdk.OsConfiguration) {

	if testutil.AssertNilEquality(test_framework, "Windows", cliOsConfiguration.Windows, sdkOsConfiguration.Windows) {
		assertEqualOsConfigurationWindows(test_framework, *cliOsConfiguration.Windows, *sdkOsConfiguration.Windows)
	}

	assert.Equal(test_framework, cliOsConfiguration.RootPassword, sdkOsConfiguration.RootPassword)
	assert.Equal(test_framework, cliOsConfiguration.ManagementUiUrl, sdkOsConfiguration.ManagementUiUrl)

	if testutil.AssertNilEquality(test_framework, "Management Access Allowed IPs", cliOsConfiguration.ManagementAccessAllowedIps, sdkOsConfiguration.ManagementAccessAllowedIps) {
		assert.Equal(test_framework, len(*cliOsConfiguration.ManagementAccessAllowedIps), len(*sdkOsConfiguration.ManagementAccessAllowedIps))

		for i := range *cliOsConfiguration.ManagementAccessAllowedIps {
			assert.Equal(test_framework, (*cliOsConfiguration.ManagementAccessAllowedIps)[i], (*sdkOsConfiguration.ManagementAccessAllowedIps)[i])
		}
	}
}
