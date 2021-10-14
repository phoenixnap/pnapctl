package server

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
}

type OsConfiguration struct {
	Windows                    *OsConfigurationWindows `yaml:"windows,omitempty" json:"windows,omitempty"`
	RootPassword               *string                 `yaml:"rootPassword,omitempty" json:"rootPassword,omitempty"`
	ManagementUiUrl            *string                 `yaml:"managementUiUrl,omitempty" json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string               `yaml:"managementAccessAllowedIps,omitempty" json:"managementAccessAllowedIps,omitempty"`
}

type OsConfigurationWindows struct {
	RdpAllowedIps *[]string `yaml:"rdpAllowedIps,omitempty" json:"rdpAllowedIps,omitempty"`
}

type TagAssignmentRequest struct {
	Name  string  `yaml:"name" json:"name"`
	Value *string `yaml:"value,omitempty" json:"value,omitempty"`
}

func CreateServerRequestFromFile() (*bmcapi.ServerCreate, error) {
	files.ExpandPath(&Filename)

	data, err := files.ReadFile(Filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverCreate ServerCreate

	err = files.Unmarshal(data, &serverCreate, commandName)

	if err != nil {
		return nil, err
	}

	return ServerCreateDtoToSdk(serverCreate), nil
}

func ServerCreateDtoToSdk(serverCreate ServerCreate) *bmcapi.ServerCreate {
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
		OsConfiguration:       osConfigurationDtoToSdk(serverCreate.OsConfiguration),
		Tags:                  tagAssignmentRequestDtoToSdk(serverCreate.Tags),
	}
}

func osConfigurationDtoToSdk(osConfiguration *OsConfiguration) *bmcapi.OsConfiguration {
	return &bmcapi.OsConfiguration{
		Windows:                    osConfigurationWindowsDtoToSdk(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func osConfigurationWindowsDtoToSdk(osConfigurationWindows *OsConfigurationWindows) *bmcapi.OsConfigurationWindows {
	return &bmcapi.OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func tagAssignmentRequestDtoToSdk(tagAssignmentRequest *[]TagAssignmentRequest) *[]bmcapi.TagAssignmentRequest {
	var list []bmcapi.TagAssignmentRequest

	for _, x := range *tagAssignmentRequest {
		converted := &bmcapi.TagAssignmentRequest{
			Name:  x.Name,
			Value: x.Value,
		}

		list = append(list, *converted)
	}

	return &list
}
