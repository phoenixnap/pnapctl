package servermodels

import (
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
)

type OsConfigurationMap struct {
	Windows *OsConfigurationWindows    `json:"windows,omitempty" yaml:"windows,omitempty"`
	Esxi    *OsConfigurationMapEsxi    `json:"esxi,omitempty" yaml:"esxi,omitempty"`
	Proxmox *OsConfigurationMapProxmox `json:"proxmox,omitempty" yaml:"proxmox,omitempty"`
}

func OsConfigurationMapToSDK(osConfMap *OsConfigurationMap) *bmcapisdk.OsConfigurationMap {
	if osConfMap == nil {
		return nil
	}

	return &bmcapisdk.OsConfigurationMap{
		Windows: osConfMap.Windows.toSdk(),
		Esxi:    osConfMap.Esxi.toSdk(),
		Proxmox: osConfMap.Proxmox.toSdk(),
	}
}
