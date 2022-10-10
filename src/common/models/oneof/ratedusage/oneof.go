package ratedusage

import "github.com/phoenixnap/go-sdk-bmc/billingapi"

/*
One of types
*/
var (
	Bandwidth       = "bandwidth"
	OperatingSystem = "operating-system"
	PublicSubnet    = "public-ip"
	Server          = "bmc-server"
	Storage         = "storage"
)

func BandwidthRecordToInner(sdk billingapi.BandwidthRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(Bandwidth)
	return billingapi.BandwidthRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func OperatingSystemRecordToInner(sdk billingapi.OperatingSystemRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(OperatingSystem)
	return billingapi.OperatingSystemRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func PublicSubnetRecordToInner(sdk billingapi.PublicSubnetRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(PublicSubnet)
	return billingapi.PublicSubnetRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func ServerRecordToInner(sdk billingapi.ServerRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(Server)
	return billingapi.ServerRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func StorageRecordToInner(sdk billingapi.StorageRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(Storage)
	return billingapi.StorageRecordAsRatedUsageGet200ResponseInner(&sdk)
}
