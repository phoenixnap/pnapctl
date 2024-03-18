package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestVolumeFromSdk(test_framework *testing.T) {
	volume := generators.Generate[networkstorageapi.Volume]()
	table := VolumeTableFromSdk(volume)

	assertVolumesEqual(test_framework, volume, table)
}

func TestShortVolumeFromSdk(test_framework *testing.T) {
	volume := generators.Generate[networkstorageapi.Volume]()
	table := ShortVolumeTableFromSdk(volume)

	assertShortVolumesEqual(test_framework, volume, table)
}

func assertVolumesEqual(test_framework *testing.T, sdkVolume networkstorageapi.Volume, tblVolume VolumeTable) {
	assert.Equal(test_framework, DerefString(sdkVolume.Id), tblVolume.Id)
	assert.Equal(test_framework, DerefString(sdkVolume.Name), tblVolume.Name)
	assert.Equal(test_framework, DerefString(sdkVolume.Description), tblVolume.Description)
	assert.Equal(test_framework, DerefString(sdkVolume.Path), tblVolume.Path)
	assert.Equal(test_framework, DerefString(sdkVolume.PathSuffix), tblVolume.PathSuffix)
	assert.Equal(test_framework, Deref(sdkVolume.CapacityInGb), tblVolume.CapacityInGb)
	assert.Equal(test_framework, DerefString(sdkVolume.Protocol), tblVolume.Protocol)
	assert.Equal(test_framework, DerefString(sdkVolume.Status), tblVolume.Status)
	assert.Equal(test_framework, DerefTimeAsString(sdkVolume.CreatedOn), tblVolume.CreatedOn)
	assert.Equal(test_framework, DerefTimeAsString(sdkVolume.DeleteRequestedOn), tblVolume.DeleteRequestedOn)
	assert.Equal(test_framework, models.PermissionsToTableString(sdkVolume.Permissions), tblVolume.Permissions)
}

func assertShortVolumesEqual(test_framework *testing.T, sdkVolume networkstorageapi.Volume, tblVolume ShortVolumeTable) {
	assert.Equal(test_framework, DerefString(sdkVolume.Id), tblVolume.Id)
	assert.Equal(test_framework, DerefString(sdkVolume.Name), tblVolume.Name)
	assert.Equal(test_framework, Deref(sdkVolume.CapacityInGb), tblVolume.CapacityInGb)
	assert.Equal(test_framework, DerefTimeAsString(sdkVolume.CreatedOn), tblVolume.CreatedOn)
}
