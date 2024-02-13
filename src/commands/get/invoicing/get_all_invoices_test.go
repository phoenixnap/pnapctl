package invoicing

import (
	"testing"

	invoicingSdk "github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getRequestParams() (string, string, string, string, int, int, string, string) {
	return Number, Status, SentOnFrom, SentOnTo, Limit, Offset, SortDirection, SortField
}

func TestGetAllInvoicesSuccess(test_framework *testing.T) {
	paginatedInvoices := generators.Generate[invoicingSdk.PaginatedInvoices]()

	// Mocking
	PrepareInvoicingMockClient(test_framework).
		InvoicesGet(getRequestParams()).
		Return(&paginatedInvoices, nil)

	ExpectToPrintSuccess(test_framework, &paginatedInvoices)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllInvoicesClientFailure(test_framework *testing.T) {
	PrepareInvoicingMockClient(test_framework).
	InvoicesGet(getRequestParams()).
		Return(nil, testutil.TestError)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllInvoicesPrinterFailure(test_framework *testing.T) {
	paginatedInvoices := generators.Generate[invoicingSdk.PaginatedInvoices]()

	PrepareInvoicingMockClient(test_framework).
		InvoicesGet(getRequestParams()).
		Return(&paginatedInvoices, nil)

	expectedErr := ExpectToPrintFailure(test_framework, &paginatedInvoices)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}