package models

import (
	"time"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type LongServer struct {
	Id                   string                `yaml:"id" json:"id"`
	Status               string                `yaml:"status" json:"status"`
	Hostname             string                `yaml:"hostname" json:"hostname"`
	Description          *string               `yaml:"description,omitempty" json:"description,omitempty"`
	Os                   string                `yaml:"os" json:"os"`
	Type                 string                `yaml:"type" json:"type"`
	Location             string                `yaml:"location" json:"location"`
	Cpu                  string                `yaml:"cpu" json:"cpu"`
	CpuCount             int32                 `yaml:"cpuCount" json:"cpuCount"`
	CoresPerCpu          int32                 `yaml:"coresPerCpu" json:"coresPerCpu"`
	CpuFrequency         float32               `yaml:"cpuFrequency" json:"cpuFrequency"`
	Ram                  string                `yaml:"ram" json:"ram"`
	Storage              string                `yaml:"storage" json:"storage"`
	PrivateIpAddresses   []string              `yaml:"privateIpAddresses" json:"privateIpAddresses"`
	PublicIpAddresses    []string              `yaml:"publicIpAddresses" json:"publicIpAddresses"`
	ReservationId        *string               `yaml:"reservationId,omitempty" json:"reservationId,omitempty"`
	PricingModel         string                `yaml:"pricingModel" json:"pricingModel"`
	Password             *string               `yaml:"password,omitempty" json:"password,omitempty"`
	NetworkType          *string               `yaml:"networkType,omitempty" json:"networkType,omitempty"`
	ClusterId            *string               `yaml:"clusterId,omitempty" json:"clusterId,omitempty"`
	Tags                 *[]TagAssignment      `yaml:"tags,omitempty" json:"tags,omitempty"`
	ProvisionedOn        *time.Time            `yaml:"provisionedOn,omitempty" json:"provisionedOn,omitempty"`
	OsConfiguration      *OsConfiguration      `yaml:"osConfiguration,omitempty" json:"osConfiguration,omitempty"`
	NetworkConfiguration *NetworkConfiguration `yaml:"networkConfiguration,omitempty" json:"networkConfiguration,omitempty"`
}

type ShortServer struct {
	ID                 string   `yaml:"id" json:"id"`
	Status             string   `yaml:"status" json:"status"`
	Name               string   `yaml:"name" json:"name"`
	Description        *string  `yaml:"description" json:"description"`
	PrivateIPAddresses []string `yaml:"privateIpAddresses" json:"privateIpAddresses"`
	PublicIPAddresses  []string `yaml:"publicIpAddresses" json:"publicIpAddresses"`
}

func ToShortServer(server bmcapi.Server) ShortServer {
	return ShortServer{
		ID:                 server.Id,
		Status:             server.Status,
		Name:               server.Hostname,
		Description:        server.Description,
		PrivateIPAddresses: server.PrivateIpAddresses,
		PublicIPAddresses:  server.PublicIpAddresses,
	}
}

func ToFullServer(server bmcapi.Server) LongServer {
	return LongServer{
		Id:                   server.Id,
		Status:               server.Status,
		Hostname:             server.Hostname,
		Description:          server.Description,
		Os:                   server.Os,
		Type:                 server.Type,
		Location:             server.Location,
		Cpu:                  server.Cpu,
		CpuCount:             server.CpuCount,
		CoresPerCpu:          server.CoresPerCpu,
		CpuFrequency:         server.CpuFrequency,
		Ram:                  server.Ram,
		Storage:              server.Storage,
		PrivateIpAddresses:   server.PrivateIpAddresses,
		PublicIpAddresses:    server.PublicIpAddresses,
		ReservationId:        server.ReservationId,
		PricingModel:         server.PricingModel,
		Password:             server.Password,
		NetworkType:          server.NetworkType,
		ClusterId:            server.ClusterId,
		Tags:                 TagAssignmentSdkToDto(server.Tags),
		ProvisionedOn:        server.ProvisionedOn,
		OsConfiguration:      OsConfigurationSdkToDto(server.OsConfiguration),
		NetworkConfiguration: NetworkConfigurationSdkToDto(&server.NetworkConfiguration),
	}
}
