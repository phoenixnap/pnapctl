package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	files "phoenixnap.com/pnapctl/common/fileprocessor"
)

type ServerReset struct {
	InstallDefaultSshKeys *bool               `json:"installDefaultSshKeys,omitempty" yaml:"installDefaultSshKeys,omitempty"`
	SshKeys               []string            `json:"sshKeys,omitempty" yaml:"sshKeys,omitempty"`
	SshKeyIds             []string            `json:"sshKeyIds,omitempty" yaml:"sshKeyIds,omitempty"`
	OsConfiguration       *OsConfigurationMap `json:"osConfiguration,omitempty" yaml:"osConfiguration,omitempty"`
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
