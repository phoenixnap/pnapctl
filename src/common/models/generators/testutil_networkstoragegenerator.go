package generators

import (
	"math/rand"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateStorageNetworkSdk() networkstorageapi.StorageNetwork {
	return networkstorageapi.StorageNetwork{
		Id:          testutil.RandSeqPointer(10),
		Name:        testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
		Status:      (*networkstorageapi.Status)(testutil.RandSeqPointer(10)),
		Location:    testutil.RandSeqPointer(10),
		NetworkId:   testutil.RandSeqPointer(10),
		Ips:         testutil.RandListStringPointer(10),
		CreatedOn:   testutil.AsPointer(time.Now()),
		Volumes:     testutil.GenN(2, GenerateVolumeSdk),
	}
}

func GenerateStorageNetworkCreateSdk() networkstorageapi.StorageNetworkCreate {
	return networkstorageapi.StorageNetworkCreate{
		Name:        testutil.RandSeq(10),
		Description: testutil.RandSeqPointer(10),
		Location:    testutil.RandSeq(10),
		Volumes:     testutil.GenN(2, GenerateVolumeCreateSdk),
	}
}

func GenerateStorageNetworkUpdateSdk() networkstorageapi.StorageNetworkUpdate {
	return networkstorageapi.StorageNetworkUpdate{
		Name:        testutil.RandSeqPointer(10),
		Description: testutil.RandSeqPointer(10),
	}
}

func GenerateVolumeSdk() networkstorageapi.Volume {
	return networkstorageapi.Volume{
		Id:           testutil.RandSeqPointer(10),
		Name:         testutil.RandSeqPointer(10),
		Description:  testutil.RandSeqPointer(10),
		Path:         testutil.RandSeqPointer(10),
		PathSuffix:   testutil.RandSeqPointer(10),
		CapacityInGb: testutil.AsPointer(rand.Int31()),
		Protocol:     testutil.RandSeqPointer(10),
		Status:       (*networkstorageapi.Status)(testutil.RandSeqPointer(10)),
		CreatedOn:    testutil.AsPointer(time.Now()),
		Permissions:  testutil.AsPointer(GeneratePermissionsSdk()),
	}
}

func GenerateVolumeCreateSdk() networkstorageapi.VolumeCreate {
	return networkstorageapi.VolumeCreate{
		Name:         testutil.RandSeq(10),
		Description:  testutil.RandSeqPointer(10),
		PathSuffix:   testutil.RandSeqPointer(10),
		CapacityInGb: rand.Int31(),
	}
}

func GeneratePermissionsSdk() networkstorageapi.Permissions {
	return networkstorageapi.Permissions{
		Nfs: testutil.AsPointer(GenerateNfsPermissionsSdk()),
	}
}

func GenerateNfsPermissionsSdk() networkstorageapi.NfsPermissions {
	return networkstorageapi.NfsPermissions{
		ReadWrite:  testutil.RandListStringPointer(10),
		ReadOnly:   testutil.RandListStringPointer(10),
		RootSquash: testutil.RandListStringPointer(10),
		NoSquash:   testutil.RandListStringPointer(10),
		AllSquash:  testutil.RandListStringPointer(10),
	}
}