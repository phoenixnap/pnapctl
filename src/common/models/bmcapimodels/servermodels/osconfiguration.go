package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

type OsConfiguration struct {
	Windows                    *OsConfigurationWindows `yaml:"windows,omitempty" json:"windows,omitempty"`
	RootPassword               *string                 `yaml:"rootPassword,omitempty" json:"rootPassword,omitempty"`
	ManagementUiUrl            *string                 `yaml:"managementUiUrl,omitempty" json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps []string                `yaml:"managementAccessAllowedIps,omitempty" json:"managementAccessAllowedIps,omitempty"`
	InstallOsToRam             *bool                   `yaml:"installOsToRam,omitempty" json:"installOsToRam,omitempty"`
}

func (osConfiguration *OsConfiguration) toSdk() *bmcapisdk.OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &bmcapisdk.OsConfiguration{
		Windows:                    osConfiguration.Windows.toSdk(),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
		InstallOsToRam:             osConfiguration.InstallOsToRam,
	}
}

func osConfigurationFromSdk(osConfiguration *bmcapisdk.OsConfiguration) *OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &OsConfiguration{
		Windows:                    osConfigurationWindowsFromSdk(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
		InstallOsToRam:             osConfiguration.InstallOsToRam,
	}
}

func OsConfigurationToTableString(osConfiguration *bmcapisdk.OsConfiguration) string {
	if osConfiguration == nil || osConfiguration.RootPassword == nil {
		return ""
	} else {
		return "Password: " + *osConfiguration.RootPassword
	}
}
