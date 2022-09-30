package ranchermodels

import (
	ranchersdk "github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
)

type SshConfig struct {
	InstallDefaultKeys *bool    `json:"installDefaultKeys" yaml:"installDefaultKeys"`
	Keys               []string `json:"keys" yaml:"keys"`
	KeyIds             []string `json:"keyIds" yaml:"keyIds"`
}

func (s *SshConfig) ToSdk() *ranchersdk.NodePoolSshConfig {
	if s == nil {
		return nil
	}

	return &ranchersdk.NodePoolSshConfig{
		InstallDefaultKeys: s.InstallDefaultKeys,
		Keys:               s.Keys,
		KeyIds:             s.KeyIds,
	}
}

func SshConfigFromSdk(config *ranchersdk.NodePoolSshConfig) *SshConfig {
	if config == nil {
		return nil
	}

	return &SshConfig{
		InstallDefaultKeys: config.InstallDefaultKeys,
		Keys:               config.Keys,
		KeyIds:             config.KeyIds,
	}
}
