package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

type RatedUsageOneOf struct {
	BandwidthRecord       *billingapi.BandwidthRecord
	OperatingSystemRecord *billingapi.OperatingSystemRecord
	PublicSubnetRecord    *billingapi.PublicSubnetRecord
	ServerRecord          *billingapi.ServerRecord
}

func RatedUsageActualFromSdk(ratedUsageOneOf billingapi.RatedUsageGet200ResponseInner) interface{} {
	if ratedUsageOneOf.BandwidthRecord != nil {
		return ratedusageoneof.BandwidthRecordFromSdk(ratedUsageOneOf.BandwidthRecord)
	}

	if ratedUsageOneOf.OperatingSystemRecord != nil {
		return ratedusageoneof.OperatingSystemRecordFromSdk(ratedUsageOneOf.OperatingSystemRecord)
	}

	if ratedUsageOneOf.PublicSubnetRecord != nil {
		return ratedusageoneof.PublicSubnetRecordFromSdk(ratedUsageOneOf.PublicSubnetRecord)
	}

	if ratedUsageOneOf.ServerRecord != nil {
		return ratedusageoneof.ServerRecordFromSdk(ratedUsageOneOf.ServerRecord)
	}

	return nil
}
