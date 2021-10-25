package models

import (
	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type OsConfiguration struct {
	Windows                    *OsConfigurationWindows `yaml:"windows,omitempty" json:"windows,omitempty"`
	RootPassword               *string                 `yaml:"rootPassword,omitempty" json:"rootPassword,omitempty"`
	ManagementUiUrl            *string                 `yaml:"managementUiUrl,omitempty" json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string               `yaml:"managementAccessAllowedIps,omitempty" json:"managementAccessAllowedIps,omitempty"`
}

type OsConfigurationWindows struct {
	RdpAllowedIps *[]string `yaml:"rdpAllowedIps,omitempty" json:"rdpAllowedIps,omitempty"`
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
