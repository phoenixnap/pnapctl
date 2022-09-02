package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type OsConfigurationWindows struct {
	RdpAllowedIps []string `yaml:"rdpAllowedIps,omitempty" json:"rdpAllowedIps,omitempty"`
}

func (osConfigurationWindows *OsConfigurationWindows) toSdk() *bmcapisdk.OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &bmcapisdk.OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func osConfigurationWindowsFromSdk(osConfigurationWindows *bmcapisdk.OsConfigurationWindows) *OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}
