package productoneof

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

type Product struct {
	ProductCommon
}

type ServerProduct struct {
	Product
	Metadata ServerProductMetadata `json:"metadata" yaml:"metadata"`
}

type ServerProductMetadata struct {
	RamInGb      float32 `json:"ramInGb" yaml:"ramInGb"`
	Cpu          string  `json:"cpu" yaml:"cpu"`
	CpuCount     float32 `json:"cpuCount" yaml:"cpuCount"`
	CoresPerCpu  float32 `json:"coresPerCpu" yaml:"coresPerCpu"`
	CpuFrequency float32 `json:"cpuFrequency" yaml:"cpuFrequency"`
	Network      string  `json:"network" yaml:"network"`
	Storage      string  `json:"storage" yaml:"storage"`
}

func ServerProductMetadataFromSdk(metadata billingapi.ServerProductMetadata) ServerProductMetadata {
	return ServerProductMetadata{
		RamInGb:      metadata.RamInGb,
		Cpu:          metadata.Cpu,
		CpuCount:     metadata.CpuCount,
		CoresPerCpu:  metadata.CoresPerCpu,
		CpuFrequency: metadata.CpuFrequency,
		Network:      metadata.Network,
		Storage:      metadata.Storage,
	}
}
