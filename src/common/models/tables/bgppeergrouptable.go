package tables

import (
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

type BgpPeerGroupTable struct {
	Id                    string   `header:"ID"`
	Status                string   `header:"Status"`
	Location              string   `header:"Location"`
	Ipv4Prefixes          []string `header:"Ipv4 Prefixes"`
	TargetAsnDetails      string   `header:"Target Asn Details"`
	ActiveAsnDetails      string   `header:"Active Asn Details"`
	Password              string   `header:"Password"`
	AdvertisedRoutes      string   `header:"Advertised Routes"`
	RpkiRoaOriginAsn      int32    `header:"RPKI ROA Origin ASN"`
	EBgpMultiHop          int32    `header:"EBGP Multi Hop"`
	PeeringLoopbacksV4    []string `header:"Peering Loopbacks V4"`
	KeepAliveTimerSeconds int32    `header:"Keep Alive Timer Seconds"`
	HoldTimerSeconds      int32    `header:"Hold Timer Seconds"`
	CreatedOn             string   `header:"Created On"`
	LastUpdatedOn         string   `header:"Last Updated On"`
}

func BgpPeerGroupFromSdk(sdk networkapi.BgpPeerGroup) BgpPeerGroupTable {
	return BgpPeerGroupTable{
		Id:                    sdk.Id,
		Status:                sdk.Status,
		Location:              sdk.Location,
		Ipv4Prefixes:          iterutils.MapRef(sdk.Ipv4Prefixes, models.BgpIpv4PrefixToTableString),
		TargetAsnDetails:      models.AsnDetailsToTableString(&sdk.TargetAsnDetails),
		ActiveAsnDetails:      models.AsnDetailsToTableString(sdk.ActiveAsnDetails),
		Password:              sdk.Password,
		AdvertisedRoutes:      sdk.AdvertisedRoutes,
		RpkiRoaOriginAsn:      int32(sdk.RpkiRoaOriginAsn),
		EBgpMultiHop:          sdk.EBgpMultiHop,
		PeeringLoopbacksV4:    sdk.PeeringLoopbacksV4,
		KeepAliveTimerSeconds: sdk.KeepAliveTimerSeconds,
		HoldTimerSeconds:      sdk.HoldTimerSeconds,
		CreatedOn:             DerefString(sdk.CreatedOn),
		LastUpdatedOn:         DerefString(sdk.LastUpdatedOn),
	}
}
