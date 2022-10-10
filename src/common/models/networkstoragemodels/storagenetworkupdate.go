package networkstoragemodels

import (
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type StorageNetworkUpdate struct {
	Name        *string `json:"name" yaml:"name"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
}

func (cli StorageNetworkUpdate) ToSdk() networkstorageapi.StorageNetworkUpdate {
	return networkstorageapi.StorageNetworkUpdate{
		Name:        cli.Name,
		Description: cli.Description,
	}
}

func CreateStorageNetworkUpdateFromFile(filename string, commandName string) (*networkstorageapi.StorageNetworkUpdate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var storageNetworkUpdate StorageNetworkUpdate

	err = files.Unmarshal(data, &storageNetworkUpdate, commandName)

	if err != nil {
		return nil, err
	}

	out := storageNetworkUpdate.ToSdk()

	return &out, nil
}
