package invoicing

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/invoicingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetInvoiceSuccess(test_framework *testing.T) {
	invoice := generators.Generate[invoicingapi.Invoice]()

	PrepareInvoicingMockClient(test_framework).
		InvoicesInvoiceIdGet(RESOURCEID).
		Return(&invoice, nil)

	ExpectToPrintSuccess(test_framework, &invoice)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetInvoiceClientFailure(test_framework *testing.T) {
	PrepareInvoicingMockClient(test_framework).
	InvoicesInvoiceIdGet(RESOURCEID).
	Return(nil, testutil.TestError)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetInvoicePrinterFailure(test_framework *testing.T) {
	invoice := generators.Generate[invoicingapi.Invoice]()

	PrepareInvoicingMockClient(test_framework).
	InvoicesInvoiceIdGet(RESOURCEID).
	Return(&invoice, nil)

	expectedErr := ExpectToPrintFailure(test_framework, &invoice)

	err := GetInvoicingCmd.RunE(GetInvoicingCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}