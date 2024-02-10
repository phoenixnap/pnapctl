package invoicing

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
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