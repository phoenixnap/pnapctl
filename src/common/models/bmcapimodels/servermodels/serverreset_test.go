package servermodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestServerResestToSdk(test_framework *testing.T) {

	sdkModel := GenerateServerResetCli()
	result := ServerResetToSDK(&sdkModel)

	assert.Equal(test_framework, result.InstallDefaultSshKeys, sdkModel.InstallDefaultSshKeys)
	assert.Equal(test_framework, result.SshKeys, sdkModel.SshKeys)
	assert.Equal(test_framework, result.SshKeyIds, sdkModel.SshKeyIds)
	testutil.AssertNilEquality(test_framework, "OsConfiguration", result.OsConfiguration, sdkModel.OsConfiguration)
}
