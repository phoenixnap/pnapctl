package quotas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
)

func TestGetQuotaSuccess(test_framework *testing.T) {

	quota := generators.GenerateQuota()
	tableQuota := tables.ToQuotaTable(quota)

	PrepareBmcApiMockClient(test_framework).
		QuotaGetById(RESOURCEID).
		Return(quota, WithResponse(200, WithBody(quota)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tableQuota, "get quotas").
		Return(nil)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetQuotaNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		QuotaGetById(RESOURCEID).
		Return(bmcapisdk.Quota{}, WithResponse(400, nil), nil)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get quotas' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetQuotaClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		QuotaGetById(RESOURCEID).
		Return(bmcapisdk.Quota{}, nil, testutil.TestError)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get quotas", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetQuotaKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		QuotaGetById(RESOURCEID).
		Return(bmcapisdk.Quota{}, nil, testutil.TestKeycloakError)

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetQuotaPrinterFailure(test_framework *testing.T) {
	quota := generators.GenerateQuota()
	tableQuota := tables.ToQuotaTable(quota)

	PrepareBmcApiMockClient(test_framework).
		QuotaGetById(RESOURCEID).
		Return(quota, WithResponse(200, WithBody(tableQuota)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tableQuota, "get quotas").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetQuotasCmd.RunE(GetQuotasCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
