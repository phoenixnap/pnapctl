package billingmodels

import (
	"math/rand"
	"time"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/billingmodels/productoneof"
	"phoenixnap.com/pnapctl/common/models/billingmodels/ratedusageoneof"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

// Rated Usage

func GenerateRatedUsageGetQueryParams() RatedUsageGetQueryParams {
	return RatedUsageGetQueryParams{
		FromYearMonth:   "2020-10",
		ToYearMonth:     "2021-10",
		ProductCategory: billingapi.BANDWIDTH.Ptr(),
	}
}

func GenerateRatedUsageMonthToDateGetQueryParams() RatedUsageMonthToDateGetQueryParams {
	return RatedUsageMonthToDateGetQueryParams{
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

type RatedUsageCommonSdk interface {
	SetId(string)
	SetProductCategory(string)
	SetProductCode(string)
	SetLocation(billingapi.LocationEnum)
	SetYearMonth(string)
	SetStartDateTime(time.Time)
	SetEndDateTime(time.Time)
	SetCost(int64)
	SetPriceModel(string)
	SetUnitPrice(float32)
	SetUnitPriceDescription(string)
	SetQuantity(float32)
	SetActive(bool)
	SetUsageSessionId(string)
	SetCorrelationId(string)
	SetReservationId(string)
}

func populateRatedUsageCommon(sdk RatedUsageCommonSdk) RatedUsageCommonSdk {
	sdk.SetId(testutil.RandSeq(10))
	sdk.SetProductCategory(string(ratedusageoneof.BANDWIDTH))
	sdk.SetProductCode(testutil.RandSeq(10))
	sdk.SetLocation("PHX")
	sdk.SetYearMonth(testutil.RandSeq(10))
	sdk.SetStartDateTime(time.Now())
	sdk.SetEndDateTime(time.Now())
	sdk.SetCost(rand.Int63())
	sdk.SetPriceModel(testutil.RandSeq(10))
	sdk.SetUnitPrice(rand.Float32())
	sdk.SetUnitPriceDescription(testutil.RandSeq(10))
	sdk.SetQuantity(rand.Float32())
	sdk.SetActive(false)
	sdk.SetUsageSessionId(testutil.RandSeq(10))
	sdk.SetCorrelationId(testutil.RandSeq(10))
	sdk.SetReservationId(testutil.RandSeq(10))
	return sdk
}

func GenerateBandwidthRecordSdk() *billingapi.BandwidthRecord {
	record := billingapi.BandwidthRecord{
		Metadata: GenerateBandwidthDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.BandwidthRecord)
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
	record := billingapi.OperatingSystemRecord{
		Metadata: GenerateOperatingSystemDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.OperatingSystemRecord)
}

func GenerateOperatingSystemDetails() billingapi.OperatingSystemDetails {
	return billingapi.OperatingSystemDetails{
		Cores:         rand.Int31(),
		CorrelationId: testutil.RandSeq(10),
	}
}

func GeneratePublicSubnetRecordSdk() *billingapi.PublicSubnetRecord {
	record := billingapi.PublicSubnetRecord{
		Metadata: GeneratePublicSubnetDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.PublicSubnetRecord)
}

func GeneratePublicSubnetDetails() billingapi.PublicSubnetDetails {
	return billingapi.PublicSubnetDetails{
		Id:   testutil.RandSeqPointer(10),
		Cidr: testutil.RandSeq(10),
		Size: testutil.RandSeq(10),
	}
}

func GenerateServerRecordSdk() *billingapi.ServerRecord {
	record := billingapi.ServerRecord{
		Metadata: GenerateServerDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.ServerRecord)
}

func GenerateServerDetails() billingapi.ServerDetails {
	return billingapi.ServerDetails{
		Id:       testutil.RandSeq(10),
		Hostname: testutil.RandSeq(10),
	}
}

// Products
func GenerateProductsGetQueryParams() ProductsGetQueryParams {
	return ProductsGetQueryParams{
		ProductCode:     testutil.RandSeqPointer(10),
		ProductCategory: testutil.RandSeqPointer(10),
		SkuCode:         testutil.RandSeqPointer(10),
		Location:        testutil.RandSeqPointer(10),
	}
}

func GenerateProduct() *billingapi.Product {
	return &billingapi.Product{
		ProductCode:     testutil.RandSeq(10),
		ProductCategory: string(productoneof.BANDWIDTH),
		Plans:           []billingapi.PricingPlan{*GeneratePricingPlan()},
	}
}

func GeneratePricingPlan() *billingapi.PricingPlan {
	return &billingapi.PricingPlan{
		Sku:                   testutil.RandSeq(10),
		SkuDescription:        testutil.RandSeqPointer(10),
		Location:              "PHX",
		PricingModel:          testutil.RandSeq(10),
		Price:                 rand.Float32(),
		PriceUnit:             billingapi.GB,
		CorrelatedProductCode: testutil.RandSeqPointer(10),
		PackageQuantity:       testutil.RanF32Pointer(),
		PackageUnit:           testutil.RandSeqPointer(10),
	}
}
