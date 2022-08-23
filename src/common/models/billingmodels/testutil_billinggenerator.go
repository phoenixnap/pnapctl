package billingmodels

import (
	"math/rand"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateGetRatedUsageQueryParams() RatedUsageGetQueryParams {
	return RatedUsageGetQueryParams{
		FromYearMonth:   "2020-10",
		ToYearMonth:     "2021-10",
		ProductCategory: billingapi.BANDWIDTH.Ptr(),
	}
}

func GenerateRatedUsageRecordSdkList() []billingapi.RatedUsageGet200ResponseInner {
	return []billingapi.RatedUsageGet200ResponseInner{
		{
			BandwidthRecord: GenerateBandwidthRecordSdk(),
		},
		{
			OperatingSystemRecord: GenerateOperatingSystemRecordSdk(),
		},
		{
			PublicSubnetRecord: GeneratePublicSubnetRecordSdk(),
		},
		{
			ServerRecord: GenerateServerRecordSdk(),
		},
	}
}

func GenerateBandwidthRecordSdk() *billingapi.BandwidthRecord {
	return &billingapi.BandwidthRecord{
		Id:                   testutil.RandSeq(10),
		ProductCategory:      ratedusageoneof.BANDWIDTH,
		ProductCode:          testutil.RandSeq(10),
		Location:             "PHX",
		YearMonth:            testutil.RandSeqPointer(10),
		StartDateTime:        time.Now(),
		EndDateTime:          time.Now(),
		Cost:                 rand.Int63(),
		PriceModel:           testutil.RandSeq(10),
		UnitPrice:            rand.Float32(),
		UnitPriceDescription: testutil.RandSeq(10),
		Quantity:             rand.Float32(),
		Active:               false,
		UsageSessionId:       testutil.RandSeq(10),
		CorrelationId:        testutil.RandSeq(10),
		ReservationId:        testutil.RandSeqPointer(10),
		Metadata:             GenerateBandwidthDetails(),
	}
}

func GenerateBandwidthDetails() billingapi.BandwidthDetails {
	return billingapi.BandwidthDetails{
		IngressGb:       rand.Float32(),
		EgressGb:        rand.Float32(),
		PackageQuantity: testutil.RanF32Pointer(),
		PackageUnit:     testutil.RandSeqPointer(10),
	}
}

func GenerateOperatingSystemRecordSdk() *billingapi.OperatingSystemRecord {
	return &billingapi.OperatingSystemRecord{
		Id:                   testutil.RandSeq(10),
		ProductCategory:      ratedusageoneof.OPERATING_SYSTEM,
		ProductCode:          testutil.RandSeq(10),
		Location:             "PHX",
		YearMonth:            testutil.RandSeqPointer(10),
		StartDateTime:        time.Now(),
		EndDateTime:          time.Now(),
		Cost:                 rand.Int63(),
		PriceModel:           testutil.RandSeq(10),
		UnitPrice:            rand.Float32(),
		UnitPriceDescription: testutil.RandSeq(10),
		Quantity:             rand.Float32(),
		Active:               false,
		UsageSessionId:       testutil.RandSeq(10),
		CorrelationId:        testutil.RandSeq(10),
		ReservationId:        testutil.RandSeqPointer(10),
		Metadata:             GenerateOperatingSystemDetails(),
	}
}

func GenerateOperatingSystemDetails() billingapi.OperatingSystemDetails {
	return billingapi.OperatingSystemDetails{
		Cores:         rand.Int31(),
		CorrelationId: testutil.RandSeq(10),
	}
}

func GeneratePublicSubnetRecordSdk() *billingapi.PublicSubnetRecord {
	return &billingapi.PublicSubnetRecord{
		Id:                   testutil.RandSeq(10),
		ProductCategory:      ratedusageoneof.PUBLIC_SUBNET,
		ProductCode:          testutil.RandSeq(10),
		Location:             "PHX",
		YearMonth:            testutil.RandSeqPointer(10),
		StartDateTime:        time.Now(),
		EndDateTime:          time.Now(),
		Cost:                 rand.Int63(),
		PriceModel:           testutil.RandSeq(10),
		UnitPrice:            rand.Float32(),
		UnitPriceDescription: testutil.RandSeq(10),
		Quantity:             rand.Float32(),
		Active:               false,
		UsageSessionId:       testutil.RandSeq(10),
		CorrelationId:        testutil.RandSeq(10),
		ReservationId:        testutil.RandSeqPointer(10),
		Metadata:             GeneratePublicSubnetDetails(),
	}
}

func GeneratePublicSubnetDetails() billingapi.PublicSubnetDetails {
	return billingapi.PublicSubnetDetails{
		Id:   testutil.RandSeqPointer(10),
		Cidr: testutil.RandSeq(10),
		Size: testutil.RandSeq(10),
	}
}

func GenerateServerRecordSdk() *billingapi.ServerRecord {
	return &billingapi.ServerRecord{
		Id:                   testutil.RandSeq(10),
		ProductCategory:      ratedusageoneof.SERVER,
		ProductCode:          testutil.RandSeq(10),
		Location:             "PHX",
		YearMonth:            testutil.RandSeqPointer(10),
		StartDateTime:        time.Now(),
		EndDateTime:          time.Now(),
		Cost:                 rand.Int63(),
		PriceModel:           testutil.RandSeq(10),
		UnitPrice:            rand.Float32(),
		UnitPriceDescription: testutil.RandSeq(10),
		Quantity:             rand.Float32(),
		Active:               false,
		UsageSessionId:       testutil.RandSeq(10),
		CorrelationId:        testutil.RandSeq(10),
		ReservationId:        testutil.RandSeqPointer(10),
		Metadata:             GenerateServerDetails(),
	}
}

func GenerateServerDetails() billingapi.ServerDetails {
	return billingapi.ServerDetails{
		Id:       testutil.RandSeq(10),
		Hostname: testutil.RandSeq(10),
	}
}
