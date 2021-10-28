package tables

import (
	"time"

	"gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/models"
)

type LongServerTable struct {
	Id                   string   `header:"ID"`
	Status               string   `header:"Status"`
	Hostname             string   `header:"Name"`
	Description          string   `header:"Description"`
	Os                   string   `header:"OS"`
	Type                 string   `header:"Type"`
	Location             string   `header:"Location"`
	Cpu                  string   `header:"Cpu"`
	CpuCount             int32    `header:"Cpu Count"`
	CoresPerCpu          int32    `header:"Cores Per Cpu"`
	CpuFrequency         float32  `header:"Cpu Frequency"`
	Ram                  string   `header:"Ram"`
	Storage              string   `header:"Storage"`
	PrivateIpAddresses   []string `header:"Private IP"`
	PublicIpAddresses    []string `header:"Public IP"`
	ReservationId        string   `header:"Reservation ID"`
	PricingModel         string   `header:"Pricing Model"`
	Password             string   `header:"Password"`
	NetworkType          string   `header:"Network Type"`
	ClusterId            string   `header:"Cluster ID"`
	Tags                 []string `header:"Tags"`
	ProvisionedOn        string   `header:"Provisioned On"`
	OsConfiguration      string   `header:"Os Configuration"`
	NetworkConfiguration string   `header:"Network Configuration"`
}

type ShortServerTable struct {
	ID                 string   `header:"id"`
	Status             string   `header:"status"`
	Name               string   `header:"name"`
	Description        string   `header:"description"`
	PrivateIPAddresses []string `header:"Private Ips"`
	PublicIPAddresses  []string `header:"Public Ips"`
}

func UpdateNillableString(str *string) *string {
	empty := ""
	if str == nil {
		return &empty
	}

	return str
}

func DerefString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func DerefTimeAsString(time *time.Time) string {
	if time == nil {
		return ""
	}

	return time.String()
}

func ToShortServerTable(server bmcapi.Server) ShortServerTable {

	return ShortServerTable{
		ID:                 server.Id,
		Status:             server.Status,
		Name:               server.Hostname,
		Description:        DerefString(server.Description),
		PrivateIPAddresses: server.PrivateIpAddresses,
		PublicIPAddresses:  server.PublicIpAddresses,
	}
}

func ToLongServerTable(server bmcapi.Server) LongServerTable {
	return LongServerTable{
		Id:                   server.Id,
		Status:               server.Status,
		Hostname:             server.Hostname,
		Description:          DerefString(server.Description),
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
		ReservationId:        DerefString(server.ReservationId),
		PricingModel:         server.PricingModel,
		Password:             DerefString(server.Password),
		NetworkType:          DerefString(server.NetworkType),
		ClusterId:            DerefString(server.ClusterId),
		Tags:                 TagsToTableStrings(server.Tags),
		ProvisionedOn:        DerefTimeAsString(server.ProvisionedOn),
		OsConfiguration:      OsConfigurationToTableString(server.OsConfiguration),
		NetworkConfiguration: NetworkConfigurationToTableString(&server.NetworkConfiguration),
	}
}

func TagsToTableStrings(tags *[]bmcapi.TagAssignment) []string {
	var tagStrings []string
	if tags == nil {
		tagStrings = []string{}
	} else {
		dtoTags := models.TagAssignmentSdkToDto(tags)
		for _, tag := range *dtoTags {
			tagStrings = append(tagStrings, tag.ToTableString())
		}
	}

	return tagStrings
}

func OsConfigurationToTableString(osConfiguration *bmcapi.OsConfiguration) string {
	if osConfiguration == nil {
		return ""
	} else {
		sdkObj := models.OsConfigurationSdkToDto(osConfiguration)
		return sdkObj.ToTableString()
	}
}

func NetworkConfigurationToTableString(networkConfiguration *bmcapi.NetworkConfiguration) string {
	if networkConfiguration == nil {
		return ""
	} else {
		sdkObj := models.NetworkConfigurationSdkToDto(networkConfiguration)
		return sdkObj.ToTableString()
	}
}
