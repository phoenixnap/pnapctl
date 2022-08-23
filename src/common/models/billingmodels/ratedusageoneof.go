package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
)

type ShortRatedUsage struct {
	Id              string
	ProductCategory string
	ProductCode     string
	YearMonth       *string
	Cost            int64
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

func ShortRatedUsageActualFromSdk(ratedUsageOneOf billingapi.RatedUsageGet200ResponseInner) interface{} {
	if ratedUsageOneOf.BandwidthRecord != nil {
		return &ShortRatedUsage{
			Id:              ratedUsageOneOf.BandwidthRecord.Id,
			ProductCategory: ratedUsageOneOf.BandwidthRecord.ProductCategory,
			ProductCode:     ratedUsageOneOf.BandwidthRecord.ProductCode,
			YearMonth:       ratedUsageOneOf.BandwidthRecord.YearMonth,
			Cost:            ratedUsageOneOf.BandwidthRecord.Cost,
		}
	}

	if ratedUsageOneOf.OperatingSystemRecord != nil {
		return &ShortRatedUsage{
			Id:              ratedUsageOneOf.OperatingSystemRecord.Id,
			ProductCategory: ratedUsageOneOf.OperatingSystemRecord.ProductCategory,
			ProductCode:     ratedUsageOneOf.OperatingSystemRecord.ProductCode,
			YearMonth:       ratedUsageOneOf.OperatingSystemRecord.YearMonth,
			Cost:            ratedUsageOneOf.OperatingSystemRecord.Cost,
		}
	}

	if ratedUsageOneOf.PublicSubnetRecord != nil {
		return &ShortRatedUsage{
			Id:              ratedUsageOneOf.PublicSubnetRecord.Id,
			ProductCategory: ratedUsageOneOf.PublicSubnetRecord.ProductCategory,
			ProductCode:     ratedUsageOneOf.PublicSubnetRecord.ProductCode,
			YearMonth:       ratedUsageOneOf.PublicSubnetRecord.YearMonth,
			Cost:            ratedUsageOneOf.PublicSubnetRecord.Cost,
		}
	}

	if ratedUsageOneOf.ServerRecord != nil {
		return &ShortRatedUsage{
			Id:              ratedUsageOneOf.ServerRecord.Id,
			ProductCategory: ratedUsageOneOf.ServerRecord.ProductCategory,
			ProductCode:     ratedUsageOneOf.ServerRecord.ProductCode,
			YearMonth:       ratedUsageOneOf.ServerRecord.YearMonth,
			Cost:            ratedUsageOneOf.ServerRecord.Cost,
		}
	}

	return nil
}
