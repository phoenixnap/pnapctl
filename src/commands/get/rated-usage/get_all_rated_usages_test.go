package rated_usage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/cmdname"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() (string, string, string) {
	return FromYearMonth, ToYearMonth, ProductCategory
}

func TestGetAllRatedUsages_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	Full = true

	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.RatedUsageRecordTableFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables).
		Return(nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsages_ShortTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	Full = false

	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.ShortRatedUsageRecordFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables).
		Return(nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsages_KeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllRatedUsages_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	var recordTables []interface{}

	for _, record := range responseList {
		recordTables = append(recordTables, *tables.ShortRatedUsageRecordFromSdk(record))
	}

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, WithResponse(200, WithBody(responseList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func TestGetAllRatedUsages_ServerError(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(nil, WithResponse(500, nil), nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}
