package quotas

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllQuotasSuccess(test_framework *testing.T) {
	quotaList := testutil.GenN(2, generators.Generate[bmcapisdk.Quota])

	var quotaTables []interface{}

	for _, quota := range quotaList {
		quotaTables = append(quotaTables, tables.ToQuotaTable(quota))
	}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quotaList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(quotaTables).
		Return(nil)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllQuotasKeycloakFailure(test_framework *testing.T) {
	quota := []bmcapisdk.Quota{generators.Generate[bmcapisdk.Quota]()}
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quota, testutil.TestKeycloakError)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllQuotasPrinterFailure(test_framework *testing.T) {
	quotaList := testutil.GenN(2, generators.Generate[bmcapisdk.Quota])

	var quotaTables []interface{}

	for _, quota := range quotaList {
		quotaTables = append(quotaTables, tables.ToQuotaTable(quota))
	}

	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quotaList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(quotaTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
