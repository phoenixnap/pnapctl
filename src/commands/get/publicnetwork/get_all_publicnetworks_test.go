package publicnetwork

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func getQueryParams() string {
	return location
}

func TestGetAllPublicNetworksSuccess(test_framework *testing.T) {
	publicNetworkList := testutil.GenN(2, generators.Generate[networkapi.PublicNetwork])

	publicNetworkTables := iterutils.MapInterface(publicNetworkList, tables.PublicNetworkTableFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(getQueryParams()).
		Return(publicNetworkList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTables).
		Return(nil)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllPublicNetworksClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(getQueryParams()).
		Return(nil, testutil.TestError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllPublicNetworksKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(getQueryParams()).
		Return(nil, testutil.TestKeycloakError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllPublicNetworksPrinterFailure(test_framework *testing.T) {
	publicNetworkList := testutil.GenN(2, generators.Generate[networkapi.PublicNetwork])
	publicNetworkTables := iterutils.MapInterface(publicNetworkList, tables.PublicNetworkTableFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksGet(getQueryParams()).
		Return(publicNetworkList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
