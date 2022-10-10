package networkstoragemodels

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// tests
func TestStorageNetworkFromSdk(test_framework *testing.T) {
	storageNetworkSdk := GenerateStorageNetworkSdk()
	storageNetworkCli := StorageNetworkFromSdk(storageNetworkSdk)

	assertEqualStorageNetwork(test_framework, storageNetworkCli, storageNetworkSdk)
}

func TestVolumeFromSdk(test_framework *testing.T) {
	volumeSdk := GenerateVolumeSdk()
	volumeCli := VolumeFromSdk(volumeSdk)

	assertEqualVolumes(test_framework, volumeCli, volumeSdk)
}

func TestPermissionsFromSdk(test_framework *testing.T) {
	permissionsSdk := GeneratePermissionsSdk()
	permissionsCli := PermissionsFromSdk(permissionsSdk)

	assertEqualPermissions(test_framework, permissionsCli, permissionsSdk)
}

// assertions
func assertEqualStorageNetwork(test_framework *testing.T, p1 StorageNetwork, p2 networkstorageapi.StorageNetwork) {
	assert.Equal(test_framework, p1.Id, p2.Id)
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.Status, p2.Status)
	assert.Equal(test_framework, p1.Location, p2.Location)
	assert.Equal(test_framework, p1.NetworkId, p2.NetworkId)
	assert.Equal(test_framework, p1.Ips, p2.Ips)
	assert.Equal(test_framework, p1.CreatedOn, p2.CreatedOn)

	testutil.ForEachPair(p1.Volumes, p2.Volumes).
		Do(test_framework, assertEqualVolumes)
}

func assertEqualVolumes(test_framework *testing.T, p1 Volume, p2 networkstorageapi.Volume) {
	assert.Equal(test_framework, p1.Id, p2.Id)
	assert.Equal(test_framework, p1.Name, p2.Name)
	assert.Equal(test_framework, p1.Description, p2.Description)
	assert.Equal(test_framework, p1.Path, p2.Path)
	assert.Equal(test_framework, p1.PathSuffix, p2.PathSuffix)
	assert.Equal(test_framework, p1.CapacityInGb, p2.CapacityInGb)
	assert.Equal(test_framework, p1.Protocol, p2.Protocol)
	assert.Equal(test_framework, p1.Status, p2.Status)
	assert.Equal(test_framework, p1.CreatedOn, p2.CreatedOn)

	assertEqualPermissions(test_framework, *p1.Permissions, *p2.Permissions)
}

func assertEqualPermissions(test_framework *testing.T, p1 Permissions, p2 networkstorageapi.Permissions) {
	assert.Equal(test_framework, p1.Nfs.ReadWrite, p2.Nfs.ReadWrite)
	assert.Equal(test_framework, p1.Nfs.ReadOnly, p2.Nfs.ReadOnly)
	assert.Equal(test_framework, p1.Nfs.RootSquash, p2.Nfs.RootSquash)
	assert.Equal(test_framework, p1.Nfs.NoSquash, p2.Nfs.NoSquash)
	assert.Equal(test_framework, p1.Nfs.AllSquash, p2.Nfs.AllSquash)
}
