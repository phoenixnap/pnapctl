package models

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type ServerCreate struct {
	Hostname              string                  `yaml:"hostname" json:"hostname"`
	Description           *string                 `yaml:"description,omitempty" json:"description,omitempty"`
	Os                    string                  `yaml:"os" json:"os"`
	Type                  string                  `yaml:"type" json:"type"`
	Location              string                  `yaml:"location" json:"location"`
	InstallDefaultSshKeys *bool                   `yaml:"installDefaultSshKeys,omitempty" json:"installDefaultSshKeys,omitempty"`
	SshKeys               *[]string               `yaml:"sshKeys,omitempty" json:"sshKeys,omitempty"`
	SshKeyIds             *[]string               `yaml:"sshKeyIds,omitempty" json:"sshKeyIds,omitempty"`
	ReservationId         *string                 `yaml:"reservationId,omitempty" json:"reservationId,omitempty"`
	PricingModel          *string                 `yaml:"pricingModel,omitempty" json:"pricingModel,omitempty"`
	NetworkType           *string                 `yaml:"networkType,omitempty" json:"networkType,omitempty"`
	OsConfiguration       *OsConfiguration        `yaml:"osConfiguration,omitempty" json:"osConfiguration,omitempty"`
	Tags                  *[]TagAssignmentRequest `yaml:"tags,omitempty" json:"tags,omitempty"`
	NetworkConfiguration  *NetworkConfiguration   `yaml:"networkConfiguration,omitempty" json:"networkConfiguration,omitempty"`
}

func CreateServerRequestFromFile(filename string, commandname string) (*bmcapi.ServerCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverCreate ServerCreate

	err = files.Unmarshal(data, &serverCreate, commandname)

	if err != nil {
		return nil, err
	}

	return serverCreate.ToSdk(), nil
}

func (serverCreate ServerCreate) ToSdk() *bmcapi.ServerCreate {
	return &bmcapi.ServerCreate{
		Hostname:              serverCreate.Hostname,
		Description:           serverCreate.Description,
		Os:                    serverCreate.Os,
		Type:                  serverCreate.Type,
		Location:              serverCreate.Location,
		InstallDefaultSshKeys: serverCreate.InstallDefaultSshKeys,
		SshKeys:               serverCreate.SshKeys,
		SshKeyIds:             serverCreate.SshKeyIds,
		ReservationId:         serverCreate.ReservationId,
		PricingModel:          serverCreate.PricingModel,
		NetworkType:           serverCreate.NetworkType,
		OsConfiguration:       serverCreate.OsConfiguration.toSdk(),
		Tags:                  mapTagAssignmentRequestToSdk(serverCreate.Tags),
		NetworkConfiguration:  serverCreate.NetworkConfiguration.toSdk(),
	}
}
