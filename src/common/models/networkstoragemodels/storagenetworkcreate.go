package networkstoragemodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type StorageNetworkCreate struct {
	Name        string         `json:"name" yaml:"name"`
	Description *string        `json:"description,omitempty" yaml:"description,omitempty"`
	Location    string         `json:"location" yaml:"location"`
	Volumes     []VolumeCreate `json:"volumes" yaml:"volumes"`
}

type VolumeCreate struct {
	Name         string  `json:"name" yaml:"name"`
	Description  *string `json:"description,omitempty" yaml:"description,omitempty"`
	PathSuffix   *string `json:"pathSuffix,omitempty" yaml:"pathSuffix,omitempty"`
	CapacityInGb int32   `json:"capacityInGb" yaml:"capacityInGb"`
}

func (cli StorageNetworkCreate) ToSdk() networkstorageapi.StorageNetworkCreate {
	return networkstorageapi.StorageNetworkCreate{
		Name:        cli.Name,
		Description: cli.Description,
		Location:    cli.Location,
		Volumes:     iterutils.Map(cli.Volumes, VolumeCreate.ToSdk),
	}
}

func (cli VolumeCreate) ToSdk() networkstorageapi.VolumeCreate {
	return networkstorageapi.VolumeCreate{
		Name:         cli.Name,
		Description:  cli.Description,
		PathSuffix:   cli.PathSuffix,
		CapacityInGb: cli.CapacityInGb,
	}
}

func CreateStorageNetworkCreateFromFile(filename string, commandName string) (*networkstorageapi.StorageNetworkCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var storageNetworkCreate StorageNetworkCreate

	err = files.Unmarshal(data, &storageNetworkCreate, commandName)

	if err != nil {
		return nil, err
	}

	out := storageNetworkCreate.ToSdk()

	return &out, nil
}
