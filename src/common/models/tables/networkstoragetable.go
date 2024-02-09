package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v2"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type StorageNetworkTable struct {
	Id                string   `header:"ID"`
	Name              string   `header:"Name"`
	Description       string   `header:"Description"`
	Status            string   `header:"Status"`
	Location          string   `header:"Location"`
	NetworkId         string   `header:"Network ID"`
	Ips               []string `header:"Ips"`
	CreatedOn         string   `header:"Created On"`
	DeleteRequestedOn string   `header:"Delete Requested On"`
	Volumes           []string `header:"Volumes"`
}

func StorageNetworkTableFromSdk(sdk networkstorageapi.StorageNetwork) StorageNetworkTable {
	return StorageNetworkTable{
		Id:                DerefString(sdk.Id),
		Name:              DerefString(sdk.Name),
		Description:       DerefString(sdk.Description),
		Status:            DerefString(sdk.Status),
		Location:          DerefString(sdk.Location),
		NetworkId:         DerefString(sdk.NetworkId),
		Ips:               sdk.Ips,
		CreatedOn:         DerefStringable(sdk.CreatedOn),
		DeleteRequestedOn: DerefStringable(sdk.DeleteRequestedOn),
		Volumes:           iterutils.MapRef(sdk.Volumes, models.VolumeToTableString),
	}
}
