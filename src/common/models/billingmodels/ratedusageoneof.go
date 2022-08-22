package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
)

type RatedUsageOneOf struct {
	BandwidthRecord       *billingapi.BandwidthRecord
	OperatingSystemRecord *billingapi.OperatingSystemRecord
	PublicSubnetRecord    *billingapi.PublicSubnetRecord
	ServerRecord          *billingapi.ServerRecord
}

type oneOfType string

const (
	BANDWIDTH        oneOfType = "BANDWIDTH RECORD"
	OPERATING_SYSTEM oneOfType = "OPERATING SYSTEM RECORD"
	PUBLIC_SUBNET    oneOfType = "PUBLIC SUBNET RECORD"
	SERVER           oneOfType = "SERVER RECORD"
	OTHER            oneOfType = "OTHER"
)

func RatedUsageActualFromSdk(ratedUsageOneOf billingapi.RatedUsageGet200ResponseInner) interface{} {
	if ratedUsageOneOf.BandwidthRecord != nil {
		return BandwidthRecordFromSdk(ratedUsageOneOf.BandwidthRecord)
	}

	if ratedUsageOneOf.OperatingSystemRecord != nil {
		return OperatingSystemRecordFromSdk(ratedUsageOneOf.OperatingSystemRecord)
	}

	if ratedUsageOneOf.PublicSubnetRecord != nil {
		return PublicSubnetRecordFromSdk(ratedUsageOneOf.PublicSubnetRecord)
	}

	if ratedUsageOneOf.ServerRecord != nil {
		return ServerRecordFromSdk(ratedUsageOneOf.ServerRecord)
	}

	return nil
}
