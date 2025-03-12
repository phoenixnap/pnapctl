package ratedusage

import "github.com/phoenixnap/go-sdk-bmc/billingapi/v3"

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
	sdk.SetProductCategory(billingapi.RATEDUSAGEPRODUCTCATEGORYENUM_BANDWIDTH)
	return billingapi.BandwidthRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func OperatingSystemRecordToInner(sdk billingapi.OperatingSystemRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(billingapi.RATEDUSAGEPRODUCTCATEGORYENUM_OPERATING_SYSTEM)
	return billingapi.OperatingSystemRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func PublicSubnetRecordToInner(sdk billingapi.PublicSubnetRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(billingapi.RATEDUSAGEPRODUCTCATEGORYENUM_PUBLIC_IP)
	return billingapi.PublicSubnetRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func ServerRecordToInner(sdk billingapi.ServerRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(billingapi.RATEDUSAGEPRODUCTCATEGORYENUM_BMC_SERVER)
	return billingapi.ServerRecordAsRatedUsageGet200ResponseInner(&sdk)
}

func StorageRecordToInner(sdk billingapi.StorageRecord) billingapi.RatedUsageGet200ResponseInner {
	sdk.SetProductCategory(billingapi.RATEDUSAGEPRODUCTCATEGORYENUM_STORAGE)
	return billingapi.StorageRecordAsRatedUsageGet200ResponseInner(&sdk)
}
