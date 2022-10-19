package productavailability

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() ([]string, []string, bool, []string, []string, float32) {
	return productCategory, productCode, showOnlyMinQuantityAvailable, location, solution, minQuantity
}

func TestGetAllProductAvailabilitiesSuccess(test_framework *testing.T) {
	productAvailabilities := testutil.GenN(5, generators.Generate[billingapi.ProductAvailability])

	productAvailabilitiesTable := iterutils.MapInterface(
		productAvailabilities,
		tables.ProductAvailabilityTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(getQueryParams()).
		Return(productAvailabilities, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(productAvailabilitiesTable).
		Return(nil)

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllProductAvailabilitiesClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(getQueryParams()).
		Return(nil, testutil.TestError)

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllProductAvailabilitiesPrinterFailure(test_framework *testing.T) {
	productAvailabilities := testutil.GenN(5, generators.Generate[billingapi.ProductAvailability])

	productAvailabilitiesTable := iterutils.MapInterface(
		productAvailabilities,
		tables.ProductAvailabilityTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(getQueryParams()).
		Return(productAvailabilities, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(productAvailabilitiesTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
