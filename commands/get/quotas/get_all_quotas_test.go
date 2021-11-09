package quotas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	bmcapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/bmcapi"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetAllQuotasSuccess(test_framework *testing.T) {
	quotaList := generators.GenerateQuotas(2)

	var quotaTables []interface{}

	for _, quota := range quotaList {
		quotaTables = append(quotaTables, tables.ToQuotaTable(quota))
	}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quotaList, WithResponse(200, WithBody(quotaList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(quotaTables, "get quotas").
		Return(nil)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllQuotasKeycloakFailure(test_framework *testing.T) {
	quota := []bmcapisdk.Quota{generators.GenerateQuota()}
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quota, nil, testutil.TestKeycloakError)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllQuotasPrinterFailure(test_framework *testing.T) {
	quotaList := generators.GenerateQuotas(2)

	var quotaTables []interface{}

	for _, quota := range quotaList {
		quotaTables = append(quotaTables, tables.ToQuotaTable(quota))
	}

	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quotaList, WithResponse(200, WithBody(quotaList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(quotaTables, "get quotas").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}