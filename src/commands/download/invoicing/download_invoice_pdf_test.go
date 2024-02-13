package invoicing

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestDownloadInvoiceSuccess(test_framework *testing.T) {
	// Mocking
	var file os.File

	ExpectSaveFileSuccess(test_framework, &file)

	PrepareInvoicingMockClient(test_framework).
		InvoicesInvoiceIdGeneratePdfPost(RESOURCEID).
		Return(&file, nil)

	// Run command
	err := DownloadInvoiceCmd.RunE(DownloadInvoiceCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDownloadInvoiceFailure(test_framework *testing.T) {
	// Setup
	var file os.File

	// Mocking
	PrepareInvoicingMockClient(test_framework).
		InvoicesInvoiceIdGeneratePdfPost(RESOURCEID).
		Return(&file, nil)

	// Mocking
	expectedErr := ExpectSaveFileFailure(test_framework, &file)

	// Run command
	err := DownloadInvoiceCmd.RunE(DownloadInvoiceCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestDownloadInvoiceClientFail(test_framework *testing.T) {

	// Mocking
	PrepareInvoicingMockClient(test_framework).
		InvoicesInvoiceIdGeneratePdfPost(RESOURCEID).
		Return(nil, testutil.TestError)

	// Run command
	err := DownloadInvoiceCmd.RunE(DownloadInvoiceCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
