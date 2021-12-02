package quotas

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	quotamodels "phoenixnap.com/pnapctl/common/models/bmcapimodels/quota"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllQuotasSuccess(test_framework *testing.T) {
	quotaList := quotamodels.GenerateQuotas(2)

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
	quota := []bmcapisdk.Quota{quotamodels.GenerateQuota()}
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		QuotasGet().
		Return(quota, nil, testutil.TestKeycloakError)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllQuotasPrinterFailure(test_framework *testing.T) {
	quotaList := quotamodels.GenerateQuotas(2)

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
