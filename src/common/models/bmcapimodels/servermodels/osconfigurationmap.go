package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type OsConfigurationMap struct {
	Windows *OsConfigurationWindows `json:"windows,omitempty" yaml:"windows,omitempty"`
	Esxi    *OsConfigurationMapEsxi `json:"esxi,omitempty" yaml:"esxi,omitempty"`
}

type OsConfigurationMapEsxi struct {
	RootPassword               *string   `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`
	ManagementUiUrl            *string   `json:"managementUiUrl,omitempty" yaml:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string `json:"managementAccessAllowedIps,omitempty" yaml:"managementAccessAllowedIps,omitempty"`
}

func OsConfigurationMapToSDK(osConfMap *OsConfigurationMap) *bmcapisdk.OsConfigurationMap {
	if osConfMap == nil {
		return nil
	}

	return &bmcapisdk.OsConfigurationMap{
		Windows: osConfMap.Windows.toSdk(),
		Esxi:    osConfMap.Esxi.toSdk(),
	}
}

func (osConfExsi *OsConfigurationMapEsxi) toSdk() *bmcapisdk.OsConfigurationMapEsxi {
	if osConfExsi == nil {
		return nil
	}

	return &bmcapisdk.OsConfigurationMapEsxi{
		RootPassword:               osConfExsi.RootPassword,
		ManagementUiUrl:            osConfExsi.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfExsi.ManagementAccessAllowedIps,
	}
}
