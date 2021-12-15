package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

type OsConfigurationMapEsxi struct {
	RootPassword               *string   `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`
	ManagementUiUrl            *string   `json:"managementUiUrl,omitempty" yaml:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string `json:"managementAccessAllowedIps,omitempty" yaml:"managementAccessAllowedIps,omitempty"`
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
