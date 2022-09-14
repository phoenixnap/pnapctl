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

// For common setting
type RatedUsageCommonSetter interface {
	SetId(string)
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

func populateRatedUsageCommon(sdk RatedUsageCommonSetter) RatedUsageCommonSetter {
	sdk.SetId(testutil.RandSeq(10))
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

// Individual oneof setting
func GenerateBandwidthRecordSdk() *billingapi.BandwidthRecord {
	record := billingapi.BandwidthRecord{
		ProductCategory: string(ratedusageoneof.BANDWIDTH),
		Metadata:        GenerateBandwidthDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.BandwidthRecord)
}

func GenerateBandwidthDetails() billingapi.BandwidthDetails {
	return billingapi.BandwidthDetails{
		IngressGb:       rand.Float32(),
		EgressGb:        rand.Float32(),
		PackageQuantity: testutil.AsPointer(rand.Float32()),
		PackageUnit:     testutil.AsPointer(testutil.RandSeq(10)),
	}
}

func GenerateOperatingSystemRecordSdk() *billingapi.OperatingSystemRecord {
	record := billingapi.OperatingSystemRecord{
		ProductCategory: string(ratedusageoneof.OPERATING_SYSTEM),
		Metadata:        GenerateOperatingSystemDetails(),
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
		ProductCategory: string(ratedusageoneof.PUBLIC_SUBNET),
		Metadata:        GeneratePublicSubnetDetails(),
	}
	return populateRatedUsageCommon(&record).(*billingapi.PublicSubnetRecord)
}

func GeneratePublicSubnetDetails() billingapi.PublicSubnetDetails {
	return billingapi.PublicSubnetDetails{
		Id:   testutil.AsPointer(testutil.RandSeq(10)),
		Cidr: testutil.RandSeq(10),
		Size: testutil.RandSeq(10),
	}
}

func GenerateServerRecordSdk() *billingapi.ServerRecord {
	record := billingapi.ServerRecord{
		ProductCategory: string(ratedusageoneof.SERVER),
		Metadata:        GenerateServerDetails(),
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
		ProductCode:     testutil.AsPointer(testutil.RandSeq(10)),
		ProductCategory: testutil.AsPointer(testutil.RandSeq(10)),
		SkuCode:         testutil.AsPointer(testutil.RandSeq(10)),
		Location:        testutil.AsPointer(testutil.RandSeq(10)),
	}
}

func GenerateProductSdkList() []billingapi.ProductsGet200ResponseInner {
	return []billingapi.ProductsGet200ResponseInner{
		{
			Product: GenerateBandwidthProduct(),
		},
		{
			Product: GenerateOperatingSystemProduct(),
		},
		{
			ServerProduct: GenerateServerProduct(),
		},
	}
}

// For common setting
type ProductCommonSetter interface {
	SetProductCode(string)
	SetPlans([]billingapi.PricingPlan)
}

func populateProductCommon(sdk ProductCommonSetter) ProductCommonSetter {
	sdk.SetProductCode(testutil.RandSeq(10))
	sdk.SetPlans(testutil.GenN(5, GeneratePricingPlan))
	return sdk
}

func GeneratePricingPlan() billingapi.PricingPlan {
	return billingapi.PricingPlan{
		Sku:                   testutil.RandSeq(10),
		SkuDescription:        testutil.AsPointer(testutil.RandSeq(10)),
		Location:              "PHX",
		PricingModel:          testutil.RandSeq(10),
		Price:                 rand.Float32(),
		PriceUnit:             billingapi.GB,
		CorrelatedProductCode: testutil.AsPointer(testutil.RandSeq(10)),
		PackageQuantity:       testutil.AsPointer(rand.Float32()),
		PackageUnit:           testutil.AsPointer(testutil.RandSeq(10)),
	}
}

// Individual oneof setting
func GenerateBandwidthProduct() *billingapi.Product {
	product := &billingapi.Product{
		ProductCategory: string(productoneof.BANDWIDTH),
	}

	return populateProductCommon(product).(*billingapi.Product)
}

func GenerateOperatingSystemProduct() *billingapi.Product {
	product := &billingapi.Product{
		ProductCategory: string(productoneof.OPERATING_SYSTEM),
	}

	return populateProductCommon(product).(*billingapi.Product)
}

func GenerateServerProduct() *billingapi.ServerProduct {
	product := &billingapi.ServerProduct{
		ProductCategory: string(productoneof.SERVER),
		Metadata:        GenerateServerProductMetadata(),
	}

	return populateProductCommon(product).(*billingapi.ServerProduct)
}

func GenerateServerProductMetadata() billingapi.ServerProductMetadata {
	return billingapi.ServerProductMetadata{
		RamInGb:      rand.Float32(),
		Cpu:          testutil.RandSeq(10),
		CpuCount:     rand.Float32(),
		CoresPerCpu:  rand.Float32(),
		CpuFrequency: rand.Float32(),
		Network:      testutil.RandSeq(10),
		Storage:      testutil.RandSeq(10),
	}
}

// Configuration Details
func GenerateProductAvailabilityGetQueryParams() *ProductAvailabilityGetQueryParams {
	return &ProductAvailabilityGetQueryParams{
		ProductCategory:              []string{"SERVER"},
		ProductCode:                  testutil.RandListStringPointer(10),
		ShowOnlyMinQuantityAvailable: true,
		Location:                     billingapi.AllowedLocationEnumEnumValues,
		Solution:                     []string{"SERVER_RANCHER"},
		MinQuantity:                  testutil.AsPointer(rand.Float32()),
	}
}

func GenerateConfigurationDetails() *billingapi.ConfigurationDetails {
	return &billingapi.ConfigurationDetails{
		ThresholdConfiguration: &billingapi.ThresholdConfigurationDetails{
			ThresholdAmount: 0.1,
		},
	}
}

// Product Availability
func GenerateProductAvailability() *billingapi.ProductAvailability {
	return &billingapi.ProductAvailability{
		ProductCode:                 testutil.RandSeq(10),
		ProductCategory:             testutil.RandSeq(10),
		LocationAvailabilityDetails: testutil.GenN(5, GenerateLocationAvailabilityDetail),
	}
}

func GenerateLocationAvailabilityDetail() billingapi.LocationAvailabilityDetail {
	return billingapi.LocationAvailabilityDetail{
		Location:             billingapi.PHX,
		MinQuantityRequested: rand.Float32(),
		MinQuantityAvailable: false,
		AvailableQuantity:    rand.Float32(),
		Solutions:            testutil.RandListStringPointer(10),
	}
}

// Reservation
func GenerateReservationGetQueryParams() *ReservationsGetQueryParams {
	return &ReservationsGetQueryParams{
		ProductCategory: billingapi.BANDWIDTH.Ptr(),
	}
}

func GenerateReservation() *billingapi.Reservation {
	return &billingapi.Reservation{
		Id:                  testutil.RandSeq(10),
		ProductCode:         testutil.RandSeq(10),
		ProductCategory:     testutil.RandSeq(10),
		Location:            billingapi.ASH,
		ReservationModel:    billingapi.FREE_TIER,
		InitialInvoiceModel: billingapi.CALENDAR_MONTH.Ptr(),
		StartDateTime:       time.Now(),
		EndDateTime:         testutil.AsPointer(time.Now()),
		LastRenewalDateTime: testutil.AsPointer(time.Now()),
		NextRenewalDateTime: testutil.AsPointer(time.Now()),
		AutoRenew:           false,
		Sku:                 testutil.RandSeq(10),
		Price:               rand.Float32(),
		PriceUnit:           billingapi.GB,
		AssignedResourceId:  testutil.AsPointer(testutil.RandSeq(10)),
	}
}

func GenerateReservationAutoRenewDisableRequestCli() *ReservationAutoRenewDisableRequest {
	return &ReservationAutoRenewDisableRequest{
		AutoRenewDisableReason: testutil.AsPointer(testutil.RandSeq(10)),
	}
}

func GenerateReservationRequestCli() *ReservationRequest {
	return &ReservationRequest{
		Sku: testutil.RandSeq(10),
	}
}
