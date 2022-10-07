package accountbillingconfiguration

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAccountBillingConfigurationSuccess(test_framework *testing.T) {
	configurationDetail := generators.GenerateConfigurationDetails()

	configurationDetailTable := tables.ConfigurationDetailsTableFromSdk(configurationDetail)

	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(&configurationDetail, WithResponse(200, WithBody(configurationDetail)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(configurationDetailTable, "get account-billing-configuration").
		Return(nil)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAccountBillingConfigurationClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(nil, WithResponse(400, nil), testutil.TestError)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get account-billing-configuration", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAccountBillingConfigurationKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAccountBillingConfigurationPrinterFailure(test_framework *testing.T) {
	configurationDetail := generators.GenerateConfigurationDetails()

	configurationDetailTable := tables.ConfigurationDetailsTableFromSdk(configurationDetail)

	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(&configurationDetail, WithResponse(200, WithBody(configurationDetail)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(configurationDetailTable, "get account-billing-configuration").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
