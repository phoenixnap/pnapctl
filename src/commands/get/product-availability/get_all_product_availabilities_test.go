package productavailability

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/queryparams/billing"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllProductAvailabilitiesSuccess(test_framework *testing.T) {
	productAvailabilities := testutil.GenN(5, generators.Generate[billingapi.ProductAvailability])
	queryParams := generators.GenerateProductAvailabilityGetQueryParams()
	setQueryParams(queryParams)

	productAvailabilitiesTable := iterutils.MapInterface(
		productAvailabilities,
		tables.ProductAvailabilityTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(queryParams).
		Return(productAvailabilities, WithResponse(200, WithBody(productAvailabilities)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(productAvailabilitiesTable).
		Return(nil)

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllProductAvailabilitiesClientFailure(test_framework *testing.T) {
	queryParams := generators.GenerateProductAvailabilityGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(queryParams).
		Return(nil, WithResponse(400, nil), testutil.TestError)

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllProductAvailabilitiesKeycloakFailure(test_framework *testing.T) {
	queryParams := generators.GenerateProductAvailabilityGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllProductAvailabilitiesPrinterFailure(test_framework *testing.T) {
	productAvailabilities := testutil.GenN(5, generators.Generate[billingapi.ProductAvailability])
	queryParams := generators.GenerateProductAvailabilityGetQueryParams()
	setQueryParams(queryParams)

	productAvailabilitiesTable := iterutils.MapInterface(
		productAvailabilities,
		tables.ProductAvailabilityTableFromSdk,
	)

	// Mocking
	PrepareBillingMockClient(test_framework).
		ProductAvailabilityGet(queryParams).
		Return(productAvailabilities, WithResponse(200, WithBody(productAvailabilities)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(productAvailabilitiesTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetProductAvailabilitiesCmd.RunE(GetProductAvailabilitiesCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func setQueryParams(queryparams billing.ProductAvailabilityGetQueryParams) {
	productCategory = queryparams.ProductCategory
	productCode = queryparams.ProductCode
	showOnlyMinQuantityAvailable = queryparams.ShowOnlyMinQuantityAvailable
	location = testutil.AsStrings(queryparams.Location)
	solution = queryparams.Solution
	minQuantity = *queryparams.MinQuantity
}
