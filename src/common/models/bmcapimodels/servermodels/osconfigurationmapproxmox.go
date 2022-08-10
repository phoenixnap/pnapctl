package servermodels

import "github.com/phoenixnap/go-sdk-bmc/bmcapi"

type OsConfigurationMapProxmox struct {
	RootPassword               *string  `json:"rootPassword,omitempty" yaml:"rootPassword,omitempty"`
	ManagementUiUrl            *string  `json:"managementUiUrl,omitempty" yaml:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps []string `json:"managementAccessAllowedIps,omitempty" yaml:"managementAccessAllowedIps,omitempty"`
}

func (osConfigurationMapProxmox *OsConfigurationMapProxmox) toSdk() *bmcapi.OsConfigurationMapProxmox {
	if osConfigurationMapProxmox == nil {
		return nil
	}

	var osConfigurationMapProxmoxSdk = bmcapi.OsConfigurationMapProxmox{
		RootPassword:               osConfigurationMapProxmox.RootPassword,
		ManagementUiUrl:            osConfigurationMapProxmox.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfigurationMapProxmox.ManagementAccessAllowedIps,
	}

	return &osConfigurationMapProxmoxSdk
}
