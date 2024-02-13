package invoicing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"phoenixnap.com/pnapctl/common/ctlerrors"
)

func TestPayInvoiceSuccess(test_framework *testing.T) {

	PrepareInvoicingMockClient(test_framework).
	InvoicesInvoiceIdPayPost(RESOURCEID).
	Return(nil, nil)

	// Run command
	err := PayInvoiceCmd.RunE(PayInvoiceCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDownloadInvoiceClientFail(test_framework *testing.T) {

	// Mocking
	PrepareInvoicingMockClient(test_framework).
	InvoicesInvoiceIdPayPost(RESOURCEID).
	Return(nil, testutil.TestError)

	// Run command
	err := PayInvoiceCmd.RunE(PayInvoiceCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}