package networkstoragemodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestStorageNetworkCreateToSdkSuccess(test_framework *testing.T) {
	cli := GenerateStorageNetworkCreateCli()
	sdk := cli.ToSdk()

	assertStorageNetworkCreateEqual(test_framework, sdk, cli)
}

func assertStorageNetworkCreateEqual(test_framework *testing.T, sdk networkstorageapi.StorageNetworkCreate, cli StorageNetworkCreate) {
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Description, cli.Description)
	assert.Equal(test_framework, sdk.Location, cli.Location)

	testutil.ForEachPair(sdk.Volumes, cli.Volumes).
		Do(test_framework, assertVolumeCreateEqual)
}

func assertVolumeCreateEqual(test_framework *testing.T, sdk networkstorageapi.VolumeCreate, cli VolumeCreate) {
	assert.Equal(test_framework, sdk.Name, cli.Name)
	assert.Equal(test_framework, sdk.Description, cli.Description)
	assert.Equal(test_framework, sdk.PathSuffix, cli.PathSuffix)
	assert.Equal(test_framework, sdk.CapacityInGb, cli.CapacityInGb)
}
