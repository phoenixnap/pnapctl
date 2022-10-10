package servermodels

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/stretchr/testify/assert"
)

// tests
func TestOsConfigurationWindowsToSdk(test_framework *testing.T) {
	cliModel := GenerateOsConfigurationWindowsCli()
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationWindows(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfiguraitonWindowsToSdk(test_framework *testing.T) {
	var cliModel *OsConfigurationWindows = nil

	assert.Nil(test_framework, cliModel.toSdk())
}

func TestOsConfigurationWindowsNilRdpAllowedIpsToSdk(test_framework *testing.T) {
	var cliModel *OsConfigurationWindows = &OsConfigurationWindows{
		RdpAllowedIps: nil,
	}
	sdkModel := cliModel.toSdk()

	assertEqualOsConfigurationWindows(test_framework, *cliModel, *sdkModel)
}

func TestOsConfigurationWindowsFromSdk(test_framework *testing.T) {
	sdkModel := GenerateOsConfigurationWindowsSdk()
	cliModel := osConfigurationWindowsFromSdk(sdkModel)

	assertEqualOsConfigurationWindows(test_framework, *cliModel, *sdkModel)
}

func TestNilOsConfiguraitonWindowsFromSdk(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfigurationWindows = nil

	assert.Nil(test_framework, osConfigurationWindowsFromSdk(sdkModel))
}

func TestOsConfigurationWindowsFromSdkNilRdpAllowedIps(test_framework *testing.T) {
	var sdkModel *bmcapisdk.OsConfigurationWindows = &bmcapisdk.OsConfigurationWindows{
		RdpAllowedIps: nil,
	}
	cliModel := osConfigurationWindowsFromSdk(sdkModel)

	assertEqualOsConfigurationWindows(test_framework, *cliModel, *sdkModel)
}

// assertion functions
func assertEqualOsConfigurationWindows(test_framework *testing.T, cliOsConfigurationWindows OsConfigurationWindows, sdkOsConfigurationWindows bmcapisdk.OsConfigurationWindows) {

	if testutil.AssertNilEquality(test_framework, "RDP Allowed IPs", cliOsConfigurationWindows.RdpAllowedIps, sdkOsConfigurationWindows.RdpAllowedIps) {
		assert.Equal(test_framework, len(cliOsConfigurationWindows.RdpAllowedIps), len(sdkOsConfigurationWindows.RdpAllowedIps))

		for i := range cliOsConfigurationWindows.RdpAllowedIps {
			assert.Equal(test_framework, (cliOsConfigurationWindows.RdpAllowedIps)[i], (sdkOsConfigurationWindows.RdpAllowedIps)[i])
		}
	}
}
