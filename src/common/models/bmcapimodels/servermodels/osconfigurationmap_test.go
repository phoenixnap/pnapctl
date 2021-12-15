package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestOsConfigurationMapToSdk(test_framework *testing.T) {
	cliModel := GenerateOsConfigurationMapCli()
	sdkModel := OsConfigurationMapToSDK(cliModel)

	assertEqualOsConfigurationMap(test_framework, *cliModel, *sdkModel)
}

func TestEmptyOsConfigurationMapToSdk(test_framework *testing.T) {
	cliModel := &OsConfigurationMap{}
	sdkModel := OsConfigurationMapToSDK(cliModel)

	assertEqualOsConfigurationMap(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfigurationMapToSdk(test_framework *testing.T) {
	var cliModel *OsConfigurationMap = nil

	assert.Nil(test_framework, OsConfigurationMapToSDK(cliModel))
}

// assertion functions
func assertEqualOsConfigurationMap(test_framework *testing.T, cliOsConfigurationMap OsConfigurationMap, sdkOsConfigurationMap bmcapisdk.OsConfigurationMap) {

	if testutil.AssertNilEquality(test_framework, "Windows Configuration", cliOsConfigurationMap.Windows, sdkOsConfigurationMap.Windows) {
		assertEqualOsConfigurationWindows(test_framework, *cliOsConfigurationMap.Windows, *sdkOsConfigurationMap.Windows)
	}

	if testutil.AssertNilEquality(test_framework, "Esxi Configuration", cliOsConfigurationMap.Esxi, sdkOsConfigurationMap.Esxi) {
		assertEqualOsConfigurationMapEsxi(test_framework, *cliOsConfigurationMap.Esxi, *sdkOsConfigurationMap.Esxi)
	}
}
