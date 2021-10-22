package server

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type ServerReset struct {
	InstallDefaultSshKeys *bool               `json:"installDefaultSshKeys,omitempty"`
	SshKeys               *[]string           `json:"sshKeys,omitempty"`
	SshKeyIds             *[]string           `json:"sshKeyIds,omitempty"`
	OsConfiguration       *OsConfigurationMap `json:"osConfiguration,omitempty"`
}

type OsConfigurationMap struct {
	Windows *OsConfigurationWindows `json:"windows,omitempty"`
	Esxi    *OsConfigurationMapEsxi `json:"esxi,omitempty"`
}

type OsConfigurationWindows struct {
	RdpAllowedIps *[]string `json:"rdpAllowedIps,omitempty"`
}

type OsConfigurationMapEsxi struct {
	RootPassword               *string   `json:"rootPassword,omitempty"`
	ManagementUiUrl            *string   `json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string `json:"managementAccessAllowedIps,omitempty"`
}

func CreateResetRequestFromFile() (*bmcapi.ServerReset, error) {
	files.ExpandPath(&Filename)
	data, err := files.ReadFile(Filename, commandName)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverReset ServerReset

	err = files.Unmarshal(data, &serverReset, commandName)

	if err != nil {
		return nil, err
	}

	return ServerResetToSDK(&serverReset), nil
}

func ServerResetToSDK(resetRequest *ServerReset) *bmcapi.ServerReset {
	if resetRequest == nil {
		return nil
	}

	return &bmcapi.ServerReset{
		InstallDefaultSshKeys: resetRequest.InstallDefaultSshKeys,
		SshKeys:               resetRequest.SshKeys,
		SshKeyIds:             resetRequest.SshKeyIds,
		OsConfiguration:       OsConfigurationMapToSDK(resetRequest.OsConfiguration),
	}
}

func OsConfigurationMapToSDK(osConfMap *OsConfigurationMap) *bmcapi.OsConfigurationMap {
	if osConfMap == nil {
		return nil
	}

	return &bmcapi.OsConfigurationMap{
		Windows: OsConfigurationWindowsToSDK(osConfMap.Windows),
		Esxi:    OsConfigurationMapEsxiToSDK(osConfMap.Esxi),
	}
}

func OsConfigurationWindowsToSDK(osConfWin *OsConfigurationWindows) *bmcapi.OsConfigurationWindows {
	if osConfWin == nil {
		return nil
	}

	return &bmcapi.OsConfigurationWindows{
		RdpAllowedIps: osConfWin.RdpAllowedIps,
	}
}

func OsConfigurationMapEsxiToSDK(osConfExsi *OsConfigurationMapEsxi) *bmcapi.OsConfigurationMapEsxi {
	if osConfExsi == nil {
		return nil
	}

	return &bmcapi.OsConfigurationMapEsxi{
		RootPassword:               osConfExsi.RootPassword,
		ManagementUiUrl:            osConfExsi.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfExsi.ManagementAccessAllowedIps,
	}
}
