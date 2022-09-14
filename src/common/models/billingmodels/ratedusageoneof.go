package billingmodels

import (
	billingapisdk "github.com/phoenixnap/go-sdk-bmc/billingapi"
	. "phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

type ShortRatedUsage struct {
	Id              string        `json:"id" yaml:"id"`
	ProductCategory Discriminator `json:"productCategory" yaml:"productCategory"`
	ProductCode     string        `json:"productCode" yaml:"productCode"`
	YearMonth       string        `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	Cost            int64         `json:"cost" yaml:"cost"`
}

func RatedUsageActualFromSdk(ratedUsageOneOf billingapisdk.RatedUsageGet200ResponseInner) interface{} {
	ratedUsage := RatedUsageFromSdkOneOf(&ratedUsageOneOf)

	if ratedUsage == nil {
		return nil
	}

	switch {
	case ratedUsage.IsActually(BANDWIDTH):
		return &BandwidthRecord{
			RatedUsageCommon: *ratedUsage,
			Metadata:         *BandwidthDetailsFromSdk(&ratedUsageOneOf.BandwidthRecord.Metadata),
		}
	case ratedUsage.IsActually(OPERATING_SYSTEM):
		return &OperatingSystemRecord{
			RatedUsageCommon: *ratedUsage,
			Metadata:         *OperatingSystemDetailsFromSdk(&ratedUsageOneOf.OperatingSystemRecord.Metadata),
		}
	case ratedUsage.IsActually(PUBLIC_SUBNET):
		return &PublicSubnetRecord{
			RatedUsageCommon: *ratedUsage,
			Metadata:         *PublicSubnetDetailsFromSdk(&ratedUsageOneOf.PublicSubnetRecord.Metadata),
		}
	case ratedUsage.IsActually(SERVER):
		return &ServerRecord{
			RatedUsageCommon: *ratedUsage,
			Metadata:         *ServerDetailsFromSdk(&ratedUsageOneOf.ServerRecord.Metadata),
		}
	}

	return nil
}

func ShortRatedUsageActualFromSdk(ratedUsageOneOf billingapisdk.RatedUsageGet200ResponseInner) interface{} {
	ratedUsage := RatedUsageFromSdkOneOf(&ratedUsageOneOf)

	if ratedUsage == nil {
		return nil
	}

	return &ShortRatedUsage{
		Id:              ratedUsage.Id,
		ProductCategory: ratedUsage.ProductCategory,
		ProductCode:     ratedUsage.ProductCode,
		YearMonth:       ratedUsage.YearMonth,
		Cost:            ratedUsage.Cost,
	}
}
