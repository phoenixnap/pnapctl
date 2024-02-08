package transactions

// import (
// 	"testing"

// 	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
// 	"github.com/stretchr/testify/assert"
// 	"phoenixnap.com/pnapctl/common/ctlerrors"
// 	"phoenixnap.com/pnapctl/common/models/generators"
// 	"phoenixnap.com/pnapctl/common/models/tables"
// 	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
// 	"phoenixnap.com/pnapctl/testsupport/testutil"
// )

// func TestGetQuotaSuccess(test_framework *testing.T) {
// 	quota := generators.Generate[bmcapi.Quota]()
// 	tableQuota := tables.ToQuotaTable(quota)

// 	PrepareBmcApiMockClient(test_framework).
// 		QuotaGetById(RESOURCEID).
// 		Return(&quota, nil)

// 	ExpectToPrintSuccess(test_framework, tableQuota)

// 	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

// 	// Assertions
// 	assert.NoError(test_framework, err)
// }

// func TestGetQuotaClientFailure(test_framework *testing.T) {
// 	PrepareBmcApiMockClient(test_framework).
// 		QuotaGetById(RESOURCEID).
// 		Return(nil, testutil.TestError)

// 	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

// 	// Expected error
// 	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

// 	// Assertions
// 	assert.EqualError(test_framework, err, expectedErr.Error())
// }

// func TestGetQuotaPrinterFailure(test_framework *testing.T) {
// 	quota := generators.Generate[bmcapi.Quota]()
// 	tableQuota := tables.ToQuotaTable(quota)

// 	PrepareBmcApiMockClient(test_framework).
// 		QuotaGetById(RESOURCEID).
// 		Return(&quota, nil)

// 	expectedErr := ExpectToPrintFailure(test_framework, tableQuota)

// 	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

// 	// Assertions
// 	assert.EqualError(test_framework, err, expectedErr.Error())
// }
