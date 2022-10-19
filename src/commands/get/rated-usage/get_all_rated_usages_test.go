package rated_usage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func getQueryParams() (string, string, string) {
	return FromYearMonth, ToYearMonth, ProductCategory
}

func TestGetAllRatedUsages_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.RatedUsageRecordTableFromSdk))
	Full = true

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, nil)

	ExpectToPrintSuccess(test_framework, recordTables)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsages_ShortTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ShortRatedUsageRecordFromSdk))
	Full = false

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, nil)

	ExpectToPrintSuccess(test_framework, recordTables)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsages_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ShortRatedUsageRecordFromSdk))

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageGet(getQueryParams()).
		Return(responseList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, recordTables)

	err := GetRatedUsageCmd.RunE(GetRatedUsageCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
