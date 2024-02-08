package transactions

// import (
// 	"testing"

// 	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
// 	"github.com/stretchr/testify/assert"
// 	"phoenixnap.com/pnapctl/common/models/generators"
// 	"phoenixnap.com/pnapctl/common/models/tables"
// 	"phoenixnap.com/pnapctl/common/utils/iterutils"
// 	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
// 	"phoenixnap.com/pnapctl/testsupport/testutil"
// )

// func TestGetAllQuotasSuccess(test_framework *testing.T) {
// 	quotaList := testutil.GenN(2, generators.Generate[bmcapisdk.Quota])
// 	quotaTables := iterutils.MapInterface(quotaList, tables.ToQuotaTable)

// 	// Mocking
// 	PrepareBmcApiMockClient(test_framework).
// 		QuotasGet().
// 		Return(quotaList, nil)

// 	ExpectToPrintSuccess(test_framework, quotaTables)

// 	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

// 	// Assertions
// 	assert.NoError(test_framework, err)
// }

// func TestGetAllQuotasPrinterFailure(test_framework *testing.T) {
// 	quotaList := testutil.GenN(2, generators.Generate[bmcapisdk.Quota])
// 	quotaTables := iterutils.MapInterface(quotaList, tables.ToQuotaTable)

// 	PrepareBmcApiMockClient(test_framework).
// 		QuotasGet().
// 		Return(quotaList, nil)

// 	expectedErr := ExpectToPrintFailure(test_framework, quotaTables)

// 	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

// 	// Assertions
// 	assert.EqualError(test_framework, err, expectedErr.Error())
// }
