package transactions

import (
	"testing"

	paymentsSdk "github.com/phoenixnap/go-sdk-bmc/paymentsapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func getRequestParams() (int, int, string, string, string, string) {
	return Limit, Offset, SortDirection, SortField, From, To
}

func TestGetAllTransactionsSuccess(test_framework *testing.T) {
	paginatedTransactions := generators.Generate[paymentsSdk.PaginatedTransactions]()

	// Mocking
	PreparePaymentsApiMockClient(test_framework).
		TransactionsGet(getRequestParams()).
		Return(&paginatedTransactions, nil)

	ExpectToPrintSuccess(test_framework, &paginatedTransactions)

	err := GetTransactionsCmd.RunE(GetTransactionsCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllTransactionsPrinterFailure(test_framework *testing.T) {
	paginatedTransactions := generators.Generate[paymentsSdk.PaginatedTransactions]()

	PreparePaymentsApiMockClient(test_framework).
		TransactionsGet(getRequestParams()).
		Return(&paginatedTransactions, nil)

	expectedErr := ExpectToPrintFailure(test_framework, &paginatedTransactions)

	err := GetTransactionsCmd.RunE(GetTransactionsCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
