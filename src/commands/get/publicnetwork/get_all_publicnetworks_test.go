package publicnetwork

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllPublicNetworksSuccess(test_framework *testing.T) {
	publicNetworkList := testutil.GenN(2, networkmodels.GeneratePublicNetworkSdk)
	queryParams := networkmodels.GeneratePublicNetworksGetQueryParams()
	setQueryParams(queryParams)

	publicNetworkTables := iterutils.MapInterface(publicNetworkList, tables.PublicNetworkTableFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(queryParams).
		Return(publicNetworkList, WithResponse(200, WithBody(publicNetworkList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTables, "get public-networks").
		Return(nil)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllPublicNetworksInvalidQueryParams(test_framework *testing.T) {
	queryParams := networkmodels.GeneratePublicNetworksGetQueryParams()
	invalid := "INVALID"
	queryParams.Location = &invalid
	setQueryParams(queryParams)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, "location 'INVALID' is invalid. Allowed values are [PHX ASH SGP NLD CHI SEA AUS]")
}

func TestGetAllPublicNetworksClientFailure(test_framework *testing.T) {
	queryParams := networkmodels.GeneratePublicNetworksGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(queryParams).
		Return(nil, WithResponse(200, nil), testutil.TestError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get public-networks", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllPublicNetworksKeycloakFailure(test_framework *testing.T) {
	queryParams := networkmodels.GeneratePublicNetworksGetQueryParams()
	setQueryParams(queryParams)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(queryParams).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllPublicNetworksPrinterFailure(test_framework *testing.T) {
	publicNetworkList := testutil.GenN(2, networkmodels.GeneratePublicNetworkSdk)
	queryParams := networkmodels.GeneratePublicNetworksGetQueryParams()
	setQueryParams(queryParams)

	publicNetworkTables := iterutils.MapInterface(publicNetworkList, tables.PublicNetworkTableFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(queryParams).
		Return(publicNetworkList, WithResponse(200, WithBody(publicNetworkList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTables, "get public-networks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}

func setQueryParams(queryparams networkmodels.PublicNetworksGetQueryParams) {
	location = *queryparams.Location
}