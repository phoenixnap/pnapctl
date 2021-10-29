package models

import "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"

type OsConfigurationMap struct {
	Windows *OsConfigurationWindows `json:"windows,omitempty"`
	Esxi    *OsConfigurationMapEsxi `json:"esxi,omitempty"`
}

type OsConfigurationMapEsxi struct {
	RootPassword               *string   `json:"rootPassword,omitempty"`
	ManagementUiUrl            *string   `json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string `json:"managementAccessAllowedIps,omitempty"`
}

func OsConfigurationMapToSDK(osConfMap *OsConfigurationMap) *bmcapi.OsConfigurationMap {
	if osConfMap == nil {
		return nil
	}

	return &bmcapi.OsConfigurationMap{
		Windows: osConfMap.Windows.toSdk(),
		Esxi:    osConfMap.Esxi.toSdk(),
	}
}

func (osConfExsi *OsConfigurationMapEsxi) toSdk() *bmcapi.OsConfigurationMapEsxi {
	if osConfExsi == nil {
		return nil
	}

	return &bmcapi.OsConfigurationMapEsxi{
		RootPassword:               osConfExsi.RootPassword,
		ManagementUiUrl:            osConfExsi.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfExsi.ManagementAccessAllowedIps,
	}
}
