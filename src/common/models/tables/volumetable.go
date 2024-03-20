package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type VolumeTable struct {
	Id                string   `header:"Id"`
	Name              string   `header:"Name"`
	Description       string   `header:"Description"`
	Path              string   `header:"Path"`
	PathSuffix        string   `header:"Path Suffix"`
	CapacityInGb      string   `header:"Capacity In Gb"`
	Protocol          string   `header:"Protocol"`
	Status            string   `header:"Status"`
	CreatedOn         string   `header:"Created On"`
	DeleteRequestedOn string   `header:"Delete Requested On"`
	Permissions       string   `header:"Permissions"`
	Tags              []string `header:"Tags"`
}

type ShortVolumeTable struct {
	Id           string `header:"Id"`
	Name         string `header:"Name"`
	CapacityInGb string `header:"Capacity In Gb"`
	CreatedOn    string `header:"Created On"`
}

func VolumeTableFromSdk(sdk networkstorageapi.Volume) VolumeTable {
	return VolumeTable{
		Id:                DerefString(sdk.Id),
		Name:              DerefString(sdk.Name),
		Description:       DerefString(sdk.Description),
		Path:              DerefString(sdk.Path),
		PathSuffix:        DerefString(sdk.PathSuffix),
		CapacityInGb:      Deref(sdk.CapacityInGb),
		Protocol:          DerefString(sdk.Protocol),
		Status:            DerefString(sdk.Status),
		CreatedOn:         DerefTimeAsString(sdk.CreatedOn),
		DeleteRequestedOn: DerefTimeAsString(sdk.DeleteRequestedOn),
		Permissions:       models.PermissionsToTableString(sdk.Permissions),
		Tags:              iterutils.MapRef(sdk.Tags, models.StorageNetworkTagAssignmentToTableString),
	}
}

func ShortVolumeTableFromSdk(sdk networkstorageapi.Volume) ShortVolumeTable {
	return ShortVolumeTable{
		Id:           DerefString(sdk.Id),
		Name:         DerefString(sdk.Name),
		CapacityInGb: Deref(sdk.CapacityInGb),
		CreatedOn:    DerefTimeAsString(sdk.CreatedOn),
	}
}
