package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
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

func (osConfiguration *OsConfiguration) toSdk() *bmcapisdk.OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &bmcapisdk.OsConfiguration{
		Windows:                    osConfiguration.Windows.toSdk(),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func (osConfigurationWindows *OsConfigurationWindows) toSdk() *bmcapisdk.OsConfigurationWindows {
	if osConfigurationWindows == nil {
		return nil
	}

	return &bmcapisdk.OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func OsConfigurationFromSdk(osConfiguration *bmcapisdk.OsConfiguration) *OsConfiguration {
	if osConfiguration == nil {
		return nil
	}

	return &OsConfiguration{
		Windows:                    osConfigurationWindowsFromSdk(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
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

func (os OsConfiguration) ToTableString() string {
	if os.RootPassword == nil {
		return ""
	} else {
		return "Password: " + *os.RootPassword
	}
}

func OsConfigurationToTableString(osConfiguration *bmcapisdk.OsConfiguration) string {
	if osConfiguration == nil {
		return ""
	} else {
		cliObj := OsConfigurationFromSdk(osConfiguration)
		return cliObj.ToTableString()
	}
}
