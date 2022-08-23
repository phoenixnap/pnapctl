package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

type ShortRatedUsage struct {
	Id              string                        `json:"id" yaml:"id"`
	ProductCategory ratedusageoneof.Discriminator `json:"productCategory" yaml:"productCategory"`
	ProductCode     string                        `json:"productCode" yaml:"productCode"`
	YearMonth       string                        `json:"yearMonth,omitempty" yaml:"yearMonth,omitempty"`
	Cost            int64                         `json:"cost" yaml:"cost"`
}

func RatedUsageActualFromSdk(ratedUsageOneOf billingapi.RatedUsageGet200ResponseInner) interface{} {
	ratedUsage := ratedusageoneof.RatedUsageFromSdkOneOf(&ratedUsageOneOf)

	if ratedUsage == nil {
		return nil
	}

	switch {
	case ratedUsage.IsActually(ratedusageoneof.BANDWIDTH):
		return &ratedusageoneof.BandwidthRecord{
			RatedUsage: *ratedUsage,
			Metadata:   *ratedusageoneof.BandwidthDetailsFromSdk(&ratedUsageOneOf.BandwidthRecord.Metadata),
		}
	case ratedUsage.IsActually(ratedusageoneof.OPERATING_SYSTEM):
		return &ratedusageoneof.OperatingSystemRecord{
			RatedUsage: *ratedUsage,
			Metadata:   *ratedusageoneof.OperatingSystemDetailsFromSdk(&ratedUsageOneOf.OperatingSystemRecord.Metadata),
		}
	case ratedUsage.IsActually(ratedusageoneof.PUBLIC_SUBNET):
		return &ratedusageoneof.PublicSubnetRecord{
			RatedUsage: *ratedUsage,
			Metadata:   *ratedusageoneof.PublicSubnetDetailsFromSdk(&ratedUsageOneOf.PublicSubnetRecord.Metadata),
		}
	case ratedUsage.IsActually(ratedusageoneof.SERVER):
		return &ratedusageoneof.ServerRecord{
			RatedUsage: *ratedUsage,
			Metadata:   *ratedusageoneof.ServerDetailsFromSdk(&ratedUsageOneOf.ServerRecord.Metadata),
		}
	}

	return nil
}

func ShortRatedUsageActualFromSdk(ratedUsageOneOf billingapi.RatedUsageGet200ResponseInner) interface{} {
	ratedUsage := ratedusageoneof.RatedUsageFromSdkOneOf(&ratedUsageOneOf)

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
