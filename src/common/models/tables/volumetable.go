package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"phoenixnap.com/pnapctl/common/models"
)

type VolumeTable struct {
	Id           string `header:"Id"`
	Name         string `header:"Name"`
	Description  string `header:"Description"`
	Path         string `header:"Path"`
	PathSuffix   string `header:"Path Suffix"`
	CapacityInGb string `header:"Capacity In Gb"`
	Protocol     string `header:"Protocol"`
	Status       string `header:"Status"`
	CreatedOn    string `header:"Created On"`
	Permissions  string `header:"Permissions"`
}

type ShortVolumeTable struct {
	Id           string `header:"Id"`
	Name         string `header:"Name"`
	CapacityInGb string `header:"Capacity In Gb"`
	CreatedOn    string `header:"Created On"`
}

func VolumeTableFromSdk(sdk networkstorageapi.Volume) VolumeTable {
	return VolumeTable{
		Id:           DerefString(sdk.Id),
		Name:         DerefString(sdk.Name),
		Description:  DerefString(sdk.Description),
		Path:         DerefString(sdk.Path),
		PathSuffix:   DerefString(sdk.PathSuffix),
		CapacityInGb: Deref(sdk.CapacityInGb),
		Protocol:     DerefString(sdk.Protocol),
		Status:       DerefString(sdk.Status),
		CreatedOn:    DerefTimeAsString(sdk.CreatedOn),
		Permissions:  models.PermissionsToTableString(sdk.Permissions),
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
