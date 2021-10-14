package printer

import (
	"time"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
)

type LongServer struct {
	Id                 string           `header:"ID" yaml:"id" json:"id"`
	Status             string           `header:"Status" yaml:"status" json:"status"`
	Hostname           string           `header:"Name" yaml:"hostname" json:"hostname"`
	Description        *string          `header:"Description" yaml:"description,omitempty" json:"description,omitempty"`
	Os                 string           `header:"OS" yaml:"os" json:"os"`
	Type               string           `header:"Type" yaml:"type" json:"type"`
	Location           string           `header:"Location" yaml:"location" json:"location"`
	Cpu                string           `header:"Cpu" yaml:"cpu" json:"cpu"`
	CpuCount           int32            `header:"Cpu Count" yaml:"cpuCount" json:"cpuCount"`
	CoresPerCpu        int32            `header:"Cores Per Cpu" yaml:"coresPerCpu" json:"coresPerCpu"`
	CpuFrequency       float32          `header:"Cpu Frequency" yaml:"cpuFrequency" json:"cpuFrequency"`
	Ram                string           `header:"Ram" yaml:"ram" json:"ram"`
	Storage            string           `header:"Storage" yaml:"storage" json:"storage"`
	PrivateIpAddresses []string         `header:"Private IP" yaml:"privateIpAddresses" json:"privateIpAddresses"`
	PublicIpAddresses  []string         `header:"Public IP" yaml:"publicIpAddresses" json:"publicIpAddresses"`
	ReservationId      *string          `header:"Reservation ID" yaml:"reservationId,omitempty" json:"reservationId,omitempty"`
	PricingModel       string           `header:"Pricing Model" yaml:"pricingModel" json:"pricingModel"`
	Password           *string          `header:"Password" yaml:"password,omitempty" json:"password,omitempty"`
	NetworkType        *string          `header:"Network Type" yaml:"networkType,omitempty" json:"networkType,omitempty"`
	ClusterId          *string          `header:"Cluster ID" yaml:"clusterId,omitempty" json:"clusterId,omitempty"`
	Tags               *[]TagAssignment `header:"Tags" yaml:"tags,omitempty" json:"tags,omitempty"`
	ProvisionedOn      *time.Time       `header:"Provisioned On" yaml:"provisionedOn,omitempty" json:"provisionedOn,omitempty"`
	OsConfiguration    *OsConfiguration `header:"Os Configuration" yaml:"osConfiguration,omitempty" json:"osConfiguration,omitempty"`
}

type ShortServer struct {
	ID                 string   `yaml:"id" json:"id" header:"id"`
	Status             string   `yaml:"status" json:"status" header:"status"`
	Name               string   `yaml:"name" json:"name" header:"name"`
	Description        string   `yaml:"description" json:"description" header:"description"`
	PrivateIPAddresses []string `yaml:"privateIpAddresses" json:"privateIpAddresses" header:"Private Ips"`
	PublicIPAddresses  []string `yaml:"publicIpAddresses" json:"publicIpAddresses" header:"Public Ips"`
}

type TagAssignment struct {
	Id           string  `yaml:"id" json:"id"`
	Name         string  `yaml:"name" json:"name"`
	Value        *string `yaml:"value,omitempty" json:"value,omitempty"`
	IsBillingTag bool    `yaml:"isBillingTag" json:"isBillingTag"`
}

type OsConfiguration struct {
	Windows                    *OsConfigurationWindows `yaml:"windows,omitempty" json:"windows,omitempty"`
	RootPassword               *string                 `yaml:"rootPassword,omitempty" json:"rootPassword,omitempty"`
	ManagementUiUrl            *string                 `yaml:"managementUiUrl,omitempty" json:"managementUiUrl,omitempty"`
	ManagementAccessAllowedIps *[]string               `yaml:"managementAccessAllowedIps,omitempty" json:"managementAccessAllowedIps,omitempty"`
}

type OsConfigurationWindows struct {
	RdpAllowedIps *[]string `yaml:"rdpAllowedIps,omitempty" json:"rdpAllowedIps,omitempty"`
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
		Id:                 server.Id,
		Status:             server.Status,
		Hostname:           server.Hostname,
		Description:        server.Description,
		Os:                 server.Os,
		Type:               server.Type,
		Location:           server.Location,
		Cpu:                server.Cpu,
		CpuCount:           server.CpuCount,
		CoresPerCpu:        server.CoresPerCpu,
		CpuFrequency:       server.CpuFrequency,
		Ram:                server.Ram,
		Storage:            server.Storage,
		PrivateIpAddresses: server.PrivateIpAddresses,
		PublicIpAddresses:  server.PublicIpAddresses,
		ReservationId:      server.ReservationId,
		PricingModel:       server.PricingModel,
		Password:           server.Password,
		NetworkType:        server.NetworkType,
		ClusterId:          server.ClusterId,
		Tags:               tagAssignmentDtoToSdk(server.Tags),
		ProvisionedOn:      server.ProvisionedOn,
		OsConfiguration:    osConfigurationDtoToSdk(server.OsConfiguration),
	}
}

func PrintServerResponse(server bmcapi.Server, full bool, commandName string) error {
	if full {
		return MainPrinter.PrintOutput(ToFullServer(server), false, commandName)
	} else {
		return MainPrinter.PrintOutput(ToShortServer(server), false, commandName)
	}
}

func PrintServerListResponse(servers []bmcapi.Server, full bool, commandName string) error {
	var serverList []interface{}

	if full {
		for _, x := range servers {
			serverList = append(serverList, ToFullServer(x))
		}
	} else {
		for _, x := range servers {
			serverList = append(serverList, ToShortServer(x))
		}
	}

	return MainPrinter.PrintOutput(serverList, false, commandName)
}

func osConfigurationDtoToSdk(osConfiguration *bmcapi.OsConfiguration) *OsConfiguration {
	return &OsConfiguration{
		Windows:                    osConfigurationWindowsDtoToSdk(osConfiguration.Windows),
		RootPassword:               osConfiguration.RootPassword,
		ManagementUiUrl:            osConfiguration.ManagementUiUrl,
		ManagementAccessAllowedIps: osConfiguration.ManagementAccessAllowedIps,
	}
}

func osConfigurationWindowsDtoToSdk(osConfigurationWindows *bmcapi.OsConfigurationWindows) *OsConfigurationWindows {
	return &OsConfigurationWindows{
		RdpAllowedIps: osConfigurationWindows.RdpAllowedIps,
	}
}

func tagAssignmentDtoToSdk(tagAssignment *[]bmcapi.TagAssignment) *[]TagAssignment {
	var list []TagAssignment

	for _, x := range *tagAssignment {
		converted := &TagAssignment{
			Id:           x.Id,
			Name:         x.Name,
			Value:        x.Value,
			IsBillingTag: x.IsBillingTag,
		}

		list = append(list, *converted)
	}

	return &list
}
