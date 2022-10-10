package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
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
	cliModel := osConfigurationFromSdk(sdkModel)

	assertEqualOsConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestEmptyOsConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfiguration = &bmcapisdk.OsConfiguration{}
	cliModel := osConfigurationFromSdk(sdkModel)

	assertEqualOsConfiguration(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfiguration = nil

	assert.Nil(test_framework, osConfigurationFromSdk(sdkModel))
}

func TestOsConfigurationToTableString(test_framework *testing.T) {
	rootPass := "abc123"
	sdkModel := &bmcapisdk.OsConfiguration{
		RootPassword: &rootPass,
	}
	tableString := OsConfigurationToTableString(sdkModel)

	assert.NotNil(test_framework, tableString)
	assert.Equal(test_framework, "Password: "+rootPass, tableString)
}

func TestNilOsConfigurationToTableString(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfiguration = nil
	tableString := OsConfigurationToTableString(sdkModel)

	assert.NotNil(test_framework, tableString)
	assert.Empty(test_framework, tableString)
}

func TestNilRootPasswordOsConfigurationToTableString(test_framework *testing.T) {
	sdkModel := &bmcapisdk.OsConfiguration{
		RootPassword: nil,
	}
	tableString := OsConfigurationToTableString(sdkModel)

	assert.NotNil(test_framework, tableString)
	assert.Empty(test_framework, tableString)
}

// assertion functions
func assertEqualOsConfiguration(test_framework *testing.T, cliOsConfiguration OsConfiguration, sdkOsConfiguration bmcapisdk.OsConfiguration) {

	if testutil.AssertNilEquality(test_framework, "Windows", cliOsConfiguration.Windows, sdkOsConfiguration.Windows) {
		assertEqualOsConfigurationWindows(test_framework, *cliOsConfiguration.Windows, *sdkOsConfiguration.Windows)
	}

	assert.Equal(test_framework, cliOsConfiguration.RootPassword, sdkOsConfiguration.RootPassword)
	assert.Equal(test_framework, cliOsConfiguration.ManagementUiUrl, sdkOsConfiguration.ManagementUiUrl)
	assert.Equal(test_framework, cliOsConfiguration.InstallOsToRam, sdkOsConfiguration.InstallOsToRam)

	if testutil.AssertNilEquality(test_framework, "Management Access Allowed IPs", cliOsConfiguration.ManagementAccessAllowedIps, sdkOsConfiguration.ManagementAccessAllowedIps) {
		assert.Equal(test_framework, len(cliOsConfiguration.ManagementAccessAllowedIps), len(sdkOsConfiguration.ManagementAccessAllowedIps))

		for i := range cliOsConfiguration.ManagementAccessAllowedIps {
			assert.Equal(test_framework, (cliOsConfiguration.ManagementAccessAllowedIps)[i], (sdkOsConfiguration.ManagementAccessAllowedIps)[i])
		}
	}
}
