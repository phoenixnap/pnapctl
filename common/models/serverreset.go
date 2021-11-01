package models

import (
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type ServerReset struct {
	InstallDefaultSshKeys *bool               `json:"installDefaultSshKeys,omitempty"`
	SshKeys               *[]string           `json:"sshKeys,omitempty"`
	SshKeyIds             *[]string           `json:"sshKeyIds,omitempty"`
	OsConfiguration       *OsConfigurationMap `json:"osConfiguration,omitempty"`
}

func CreateResetRequestFromFile(filename string, commandname string) (*bmcapisdk.ServerReset, error) {
	files.ExpandPath(&filename)
	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	// Marshal file into JSON using the struct
	var serverReset ServerReset

	err = files.Unmarshal(data, &serverReset, commandname)

	if err != nil {
		return nil, err
	}

	return ServerResetToSDK(&serverReset), nil
}

func ServerResetToSDK(resetRequest *ServerReset) *bmcapisdk.ServerReset {
	if resetRequest == nil {
		return nil
	}

	return &bmcapisdk.ServerReset{
		InstallDefaultSshKeys: resetRequest.InstallDefaultSshKeys,
		SshKeys:               resetRequest.SshKeys,
		SshKeyIds:             resetRequest.SshKeyIds,
		OsConfiguration:       OsConfigurationMapToSDK(resetRequest.OsConfiguration),
	}
}
