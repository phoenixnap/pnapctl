package month_to_date

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/models/tables"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllRatedUsagesMonthToDate_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()
	setQueryParams(queryParams)

	Full = true

	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.RatedUsageRecordTableFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage month-to-date").
		Return(nil)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

// Currently the short table is an empty struct.
func TestGetAllRatedUsagesMonthToDate_ShortTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()
	setQueryParams(queryParams)

	Full = false

	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.ShortRatedUsageRecordFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage month-to-date").
		Return(nil)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsagesMonthToDate_KeycloakFailure(test_framework *testing.T) {
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllRatedUsagesMonthToDate_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()
	setQueryParams(queryParams)

	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.ShortRatedUsageRecordFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(queryParams).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables, "get rated-usage month-to-date").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetAllRatedUsagesMonthToDate_ServerError(test_framework *testing.T) {
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(queryParams).
		Return(nil, WithResponse(500, nil), nil)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	expectedMessage := "Command 'get rated-usage month-to-date' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetAllRatedUsagesMonthToDate_InvalidParams(test_framework *testing.T) {
	queryParams := generators.GenerateRatedUsageMonthToDateGetQueryParams()

	var invalidCategory = billingapi.ProductCategoryEnum("NONE")
	queryParams.ProductCategory = &invalidCategory

	setQueryParams(queryParams)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.Equal(test_framework, ctlerrors.InvalidFlagValuePassedError("category", "NONE", billingapi.AllowedProductCategoryEnumEnumValues), err)
}

func setQueryParams(queryparams billing.RatedUsageMonthToDateGetQueryParams) {
	ProductCategory = string(*queryparams.ProductCategory)
}
