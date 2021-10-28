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

func (osConfiguration *OsConfiguration) toSdk() *bmcapi.OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &bmcapi.OsConfiguration{
		Windows:                    osConfiguration.Windows.toSdk(),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func (osConfigurationWindows *OsConfigurationWindows) toSdk() *bmcapi.OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &bmcapi.OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func OsConfigurationSdkToDto(osConfiguration *bmcapi.OsConfiguration) *OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &OsConfiguration{
		Windows:                    osConfigurationWindowsSdkToDto(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func osConfigurationWindowsSdkToDto(osConfigurationWindows *bmcapi.OsConfigurationWindows) *OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func (os OsConfiguration) ToTableString() string {
	if os.RootPassword == nil {
		return ""
	} else {
		return "Password: " + *os.RootPassword
	}
}
