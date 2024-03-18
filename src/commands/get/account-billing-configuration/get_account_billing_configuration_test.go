package accountbillingconfiguration

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAccountBillingConfigurationSuccess(test_framework *testing.T) {
	configurationDetail := generators.Generate[billingapi.ConfigurationDetails]()

	configurationDetailTable := tables.ConfigurationDetailsTableFromSdk(configurationDetail)

	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(&configurationDetail, nil)

	ExpectToPrintSuccess(test_framework, configurationDetailTable)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAccountBillingConfigurationClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(nil, testutil.TestError)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAccountBillingConfigurationPrinterFailure(test_framework *testing.T) {
	configurationDetail := generators.Generate[billingapi.ConfigurationDetails]()

	configurationDetailTable := tables.ConfigurationDetailsTableFromSdk(configurationDetail)

	// Mocking
	PrepareBillingMockClient(test_framework).
		AccountBillingConfigurationGet().
		Return(&configurationDetail, nil)

	expectedErr := ExpectToPrintFailure(test_framework, configurationDetailTable)

	err := GetAccountBillingConfigurationCmd.RunE(GetAccountBillingConfigurationCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
