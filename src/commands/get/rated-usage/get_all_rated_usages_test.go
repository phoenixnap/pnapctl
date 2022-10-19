package rated_usage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
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
		Return(responseList, nil)

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
		Return(responseList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables).
		Return(nil)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
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
		Return(responseList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(recordTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
