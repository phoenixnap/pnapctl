package billingmodels

import (
	"math/rand"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateBandwidthRecordSdk() *billingapi.BandwidthRecord {
	return &billingapi.BandwidthRecord{
		Id:                   testutil.RandSeq(10),
		ProductCategory:      testutil.RandSeq(10),
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
		ProductCategory:      testutil.RandSeq(10),
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
		ProductCategory:      testutil.RandSeq(10),
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
		ProductCategory:      testutil.RandSeq(10),
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
