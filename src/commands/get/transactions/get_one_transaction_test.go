package transactions

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
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

func TestGetTransactionClientFailure(test_framework *testing.T) {
	PreparePaymentsApiMockClient(test_framework).
		TransactionGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetTransactionsCmd.RunE(GetTransactionsCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetTransactionPrinterFailure(test_framework *testing.T) {
	transaction := generators.Generate[paymentsapi.Transaction]()

	PreparePaymentsApiMockClient(test_framework).
		TransactionGetById(RESOURCEID).
		Return(&transaction, nil)

	expectedErr := ExpectToPrintFailure(test_framework, &transaction)

	err := GetTransactionsCmd.RunE(GetTransactionsCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
