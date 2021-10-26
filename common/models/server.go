package models

import (
	"time"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type LongServer struct {
	Id                   string                `header:"ID" yaml:"id" json:"id"`
	Status               string                `header:"Status" yaml:"status" json:"status"`
	Hostname             string                `header:"Name" yaml:"hostname" json:"hostname"`
	Description          *string               `header:"Description" yaml:"description,omitempty" json:"description,omitempty"`
	Os                   string                `header:"OS" yaml:"os" json:"os"`
	Type                 string                `header:"Type" yaml:"type" json:"type"`
	Location             string                `header:"Location" yaml:"location" json:"location"`
	Cpu                  string                `header:"Cpu" yaml:"cpu" json:"cpu"`
	CpuCount             int32                 `header:"Cpu Count" yaml:"cpuCount" json:"cpuCount"`
	CoresPerCpu          int32                 `header:"Cores Per Cpu" yaml:"coresPerCpu" json:"coresPerCpu"`
	CpuFrequency         float32               `header:"Cpu Frequency" yaml:"cpuFrequency" json:"cpuFrequency"`
	Ram                  string                `header:"Ram" yaml:"ram" json:"ram"`
	Storage              string                `header:"Storage" yaml:"storage" json:"storage"`
	PrivateIpAddresses   []string              `header:"Private IP" yaml:"privateIpAddresses" json:"privateIpAddresses"`
	PublicIpAddresses    []string              `header:"Public IP" yaml:"publicIpAddresses" json:"publicIpAddresses"`
	ReservationId        *string               `header:"Reservation ID" yaml:"reservationId,omitempty" json:"reservationId,omitempty"`
	PricingModel         string                `header:"Pricing Model" yaml:"pricingModel" json:"pricingModel"`
	Password             *string               `header:"Password" yaml:"password,omitempty" json:"password,omitempty"`
	NetworkType          *string               `header:"Network Type" yaml:"networkType,omitempty" json:"networkType,omitempty"`
	ClusterId            *string               `header:"Cluster ID" yaml:"clusterId,omitempty" json:"clusterId,omitempty"`
	Tags                 *[]TagAssignment      `header:"Tags" yaml:"tags,omitempty" json:"tags,omitempty"`
	ProvisionedOn        *time.Time            `header:"Provisioned On" yaml:"provisionedOn,omitempty" json:"provisionedOn,omitempty"`
	OsConfiguration      *OsConfiguration      `header:"Os Configuration" yaml:"osConfiguration,omitempty" json:"osConfiguration,omitempty"`
	NetworkConfiguration *NetworkConfiguration `header:"Network Configuration" yaml:"networkConfiguration,omitempty" json:"networkConfiguration,omitempty"`
}

type ShortServer struct {
	ID                 string   `yaml:"id" json:"id" header:"id"`
	Status             string   `yaml:"status" json:"status" header:"status"`
	Name               string   `yaml:"name" json:"name" header:"name"`
	Description        string   `yaml:"description" json:"description" header:"description"`
	PrivateIPAddresses []string `yaml:"privateIpAddresses" json:"privateIpAddresses" header:"Private Ips"`
	PublicIPAddresses  []string `yaml:"publicIpAddresses" json:"publicIpAddresses" header:"Public Ips"`
}

func ToShortServer(server bmcapi.Server) ShortServer {
	return ShortServer{
		ID:                 server.Id,
		Status:             server.Status,
		Name:               server.Hostname,
		Description:        *server.Description,
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
		Tags:                 tagAssignmentSdkToDto(server.Tags),
		ProvisionedOn:        server.ProvisionedOn,
		OsConfiguration:      osConfigurationSdkToDto(server.OsConfiguration),
		NetworkConfiguration: networkConfigurationSdkToDto(&server.NetworkConfiguration),
	}
}
