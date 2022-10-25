package month_to_date

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() string {
	return ProductCategory
}

func TestGetAllRatedUsagesMonthToDate_FullTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.RatedUsageRecordTableFromSdk))
	Full = true

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(getQueryParams()).
		Return(responseList, nil)

	ExpectToPrintSuccess(test_framework, recordTables)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

// Currently the short table is an empty struct.
func TestGetAllRatedUsagesMonthToDate_ShortTable(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ShortRatedUsageRecordFromSdk))
	Full = false

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(getQueryParams()).
		Return(responseList, nil)

	ExpectToPrintSuccess(test_framework, recordTables)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllRatedUsagesMonthToDate_ClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(getQueryParams()).
		Return(nil, testutil.TestError)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// AssertionsqueryParams
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllRatedUsagesMonthToDate_PrinterFailure(test_framework *testing.T) {
	responseList := generators.GenerateRatedUsageRecordSdkList()
	recordTables := iterutils.DerefInterface(iterutils.MapInterface(responseList, tables.ShortRatedUsageRecordFromSdk))

	// Mocking
	PrepareBillingMockClient(test_framework).
		RatedUsageMonthToDateGet(getQueryParams()).
		Return(responseList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, recordTables)

	err := GetRatedUsageMonthToDateCmd.RunE(GetRatedUsageMonthToDateCmd, []string{})

	// AssertionsqueryParams
	assert.EqualError(test_framework, err, expectedErr.Error())
}
