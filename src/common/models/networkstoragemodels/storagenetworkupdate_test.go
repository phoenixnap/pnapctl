package networkstoragemodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
)

func TestStorageNetworkUpdateToSdkSuccess(test_framework *testing.T) {
	cli := GenerateStorageNetworkUpdateCli()
	sdk := cli.ToSdk()

	assertStorageNetworkUpdateEqual(test_framework, sdk, cli)
}

func assertStorageNetworkUpdateEqual(test_framework *testing.T, sdk networkstorageapi.StorageNetworkUpdate, cli StorageNetworkUpdate) {
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Description, cli.Description)
}
