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
	NetworkConfiguration  *NetworkConfiguration   `yaml:"networkConfiguration,omitempty" json:"networkConfiguration,omitempty"`
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

type NetworkConfiguration struct {
	PrivateNetworkConfiguration *PrivateNetworkConfiguration `yaml:"privateNetworkConfiguration" json:"privateNetworkConfiguration"`
}

type PrivateNetworkConfiguration struct {
	GatewayAddress    *string                 `yaml:"gatewayAddress" json:"gatewayAddress"`
	ConfigurationType *string                 `yaml:"configurationType" json:"configurationType"`
	PrivateNetworks   *[]ServerPrivateNetwork `yaml:"privateNetworks" json:"privateNetworks"`
}

type ServerPrivateNetwork struct {
	Id                string    `yaml:"id" json:"id"`
	Ips               *[]string `yaml:"ips" json:"ips"`
	Dhcp              *bool     `yaml:"dhcp" json:"dhcp"`
	StatusDescription *string   `yaml:"statusDescription" json:"statusDescription"`
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
		NetworkConfiguration:  networkConfigurationDtoToSdk(serverCreate.NetworkConfiguration),
	}
}

func osConfigurationDtoToSdk(osConfiguration *OsConfiguration) *bmcapi.OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &bmcapi.OsConfiguration{
		Windows:                    osConfigurationWindowsDtoToSdk(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func osConfigurationWindowsDtoToSdk(osConfigurationWindows *OsConfigurationWindows) *bmcapi.OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &bmcapi.OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func tagAssignmentRequestDtoToSdk(tagAssignmentRequest *[]TagAssignmentRequest) *[]bmcapi.TagAssignmentRequest {
	if tagAssignmentRequest == nil {
		return nil
	}

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

func networkConfigurationDtoToSdk(networkConf *NetworkConfiguration) *bmcapi.NetworkConfiguration {
	if networkConf == nil {
		return nil
	}

	return &bmcapi.NetworkConfiguration{
		PrivateNetworkConfiguration: privateNetworkConfigurationDtoToSdk(networkConf.PrivateNetworkConfiguration),
	}
}

func privateNetworkConfigurationDtoToSdk(privateNetConf *PrivateNetworkConfiguration) *bmcapi.PrivateNetworkConfiguration {
	if privateNetConf == nil {
		return nil
	}

	return &bmcapi.PrivateNetworkConfiguration{
		GatewayAddress:    privateNetConf.GatewayAddress,
		ConfigurationType: privateNetConf.ConfigurationType,
		PrivateNetworks:   privateNetworksDtoToSdk(privateNetConf.PrivateNetworks),
	}
}

func privateNetworksDtoToSdk(privateNetworks *[]ServerPrivateNetwork) *[]bmcapi.ServerPrivateNetwork {
	if privateNetworks == nil {
		return nil
	}

	var bmcPrivNet []bmcapi.ServerPrivateNetwork

	for _, x := range *privateNetworks {
		bmcPrivNet = append(bmcPrivNet, bmcapi.ServerPrivateNetwork{
			Id:                x.Id,
			Ips:               x.Ips,
			Dhcp:              x.Dhcp,
			StatusDescription: x.StatusDescription,
		})
	}

	return &bmcPrivNet
}
