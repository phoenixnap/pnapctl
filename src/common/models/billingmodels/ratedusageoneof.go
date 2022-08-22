package billingmodels

import (
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/tables"
)

type Format int

const (
	SDK Format = iota
	SHORT_TABLE
	FULL_TABLE
)

type RatedUsageOneOf struct {
	BandwidthRecord       *billingapi.BandwidthRecord
	OperatingSystemRecord *billingapi.OperatingSystemRecord
	PublicSubnetRecord    *billingapi.PublicSubnetRecord
	ServerRecord          *billingapi.ServerRecord
}

func RatedUsageOneOfFromSdk(ratedUsageOneOf *billingapi.RatedUsageGet200ResponseInner) *RatedUsageOneOf {
	if ratedUsageOneOf == nil {
		return nil
	}

	return &RatedUsageOneOf{
		BandwidthRecord:       ratedUsageOneOf.BandwidthRecord,
		OperatingSystemRecord: ratedUsageOneOf.OperatingSystemRecord,
		PublicSubnetRecord:    ratedUsageOneOf.PublicSubnetRecord,
		ServerRecord:          ratedUsageOneOf.ServerRecord,
	}
}

func (oneOf *RatedUsageOneOf) GetActualInstanceAs(format Format) interface{} {
	if oneOf.BandwidthRecord != nil {
		switch format {
		case SDK:
			return BandwidthRecordFromSdk(oneOf.BandwidthRecord)
		case FULL_TABLE:
			return tables.BandwidthRecordTableFromSdk(*oneOf.BandwidthRecord)
		case SHORT_TABLE:
			return tables.ShortBandwidthRecordTableFromSdk(*oneOf.BandwidthRecord)
		}
	}

	if oneOf.OperatingSystemRecord != nil {
		switch format {
		case SDK:
			return OperatingSystemRecordFromSdk(oneOf.OperatingSystemRecord)
		case FULL_TABLE:
			return tables.OperatingSystemRecordTableFromSdk(*oneOf.OperatingSystemRecord)
		case SHORT_TABLE:
			return tables.ShortOperatingSystemRecordTableFromSdk(*oneOf.OperatingSystemRecord)
		}
	}

	if oneOf.PublicSubnetRecord != nil {
		switch format {
		case SDK:
			return PublicSubnetRecordFromSdk(oneOf.PublicSubnetRecord)
		case FULL_TABLE:
			return tables.PublicSubnetRecordTableFromSdk(*oneOf.PublicSubnetRecord)
		case SHORT_TABLE:
			return tables.ShortPublicSubnetRecordTableFromSdk(*oneOf.PublicSubnetRecord)
		}
	}

	if oneOf.ServerRecord != nil {
		switch format {
		case SDK:
			return ServerRecordFromSdk(oneOf.ServerRecord)
		case FULL_TABLE:
			return tables.ServerRecordTableFromSdk(*oneOf.ServerRecord)
		case SHORT_TABLE:
			return tables.ShortServerRecordTableFromSdk(*oneOf.ServerRecord)
		}
	}

	return nil
}
