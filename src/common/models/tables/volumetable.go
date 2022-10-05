package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"phoenixnap.com/pnapctl/common/models/networkstoragemodels"
)

type VolumeTable struct {
	Id           string
	Name         string
	Description  string
	Path         string
	PathSuffix   string
	CapacityInGb string
	Protocol     string
	Status       string
	CreatedOn    string
	Permissions  string
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
		Permissions:  networkstoragemodels.PermissionsToTableString(sdk.Permissions),
	}
}
