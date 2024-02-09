package transactions

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func TestGetTransactionSuccess(test_framework *testing.T) {
	transaction := generators.Generate[paymentsapi.Transaction]()

	PreparePaymentsApiMockClient(test_framework).
		TransactionGetById(RESOURCEID).
		Return(&transaction, nil)

	ExpectToPrintSuccess(test_framework, &transaction)

	err := GetTransactionsCmd.RunE(GetTransactionsCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

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
