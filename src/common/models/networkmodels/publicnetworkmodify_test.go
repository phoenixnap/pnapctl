package networkmodels

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/networkapi"
)

func TestPublicNetworkModifyToSdkSuccess(test_framework *testing.T) {
	cli := GeneratePublicNetworkModifyCli()
	sdk := cli.ToSdk()

	assertPublicNetworkModifyEqual(test_framework, sdk, cli)
}

func assertPublicNetworkModifyEqual(test_framework *testing.T, sdk networkapi.PublicNetworkModify, cli PublicNetworkModify) {
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Description, cli.Description)
}
