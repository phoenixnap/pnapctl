package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestStorageNetworkFromSdk(test_framework *testing.T) {
	storageNetwork := generators.Generate[networkstorageapi.StorageNetwork]()
	table := StorageNetworkTableFromSdk(storageNetwork)

	assertStorageNetworksEqual(test_framework, storageNetwork, table)
}

func assertStorageNetworksEqual(test_framework *testing.T, sdkStorageNetwork networkstorageapi.StorageNetwork, tblStorageNetwork StorageNetworkTable) {
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.Id), tblStorageNetwork.Id)
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.Name), tblStorageNetwork.Name)
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.Description), tblStorageNetwork.Description)
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.Status), tblStorageNetwork.Status)
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.Location), tblStorageNetwork.Location)
	assert.Equal(test_framework, DerefString(sdkStorageNetwork.NetworkId), tblStorageNetwork.NetworkId)
	assert.Equal(test_framework, sdkStorageNetwork.Ips, tblStorageNetwork.Ips)
	assert.Equal(test_framework, DerefStringable(sdkStorageNetwork.CreatedOn), tblStorageNetwork.CreatedOn)
	assert.Equal(test_framework, DerefStringable(sdkStorageNetwork.DeleteRequestedOn), tblStorageNetwork.DeleteRequestedOn)
	assert.Equal(test_framework, iterutils.MapRef(sdkStorageNetwork.Volumes, models.VolumeToTableString), tblStorageNetwork.Volumes)
}
