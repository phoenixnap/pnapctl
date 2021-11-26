package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi"
)

type SshConfig struct {
	InstallDefaultKeys *bool     `json:"installDefaultKeys" yaml:"installDefaultKeys"`
	Keys               *[]string `json:"keys" yaml:"keys"`
	KeyIds             *[]string `json:"keyIds" yaml:"keyIds"`
}

func (s *SshConfig) ToSdk() *ranchersdk.SshConfig {
	return &ranchersdk.SshConfig{
		InstallDefaultKeys: s.InstallDefaultKeys,
		Keys:               s.Keys,
		KeyIds:             s.KeyIds,
	}
}

func SshConfigFromSdk(config *ranchersdk.SshConfig) *SshConfig {
	return &SshConfig{
		InstallDefaultKeys: config.InstallDefaultKeys,
		Keys:               config.Keys,
		KeyIds:             config.KeyIds,
	}
}
