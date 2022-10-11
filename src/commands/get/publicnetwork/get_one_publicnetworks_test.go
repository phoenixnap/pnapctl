package publicnetwork

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetPublicNetworkSuccess(test_framework *testing.T) {
	publicNetworkSdk := generators.Generate[networkapi.PublicNetwork]()
	publicNetworkTable := tables.PublicNetworkTableFromSdk(publicNetworkSdk)

	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(&publicNetworkSdk, WithResponse(200, WithBody(publicNetworkSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTable).
		Return(nil)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

// GetPublicNetworksCmd.SetArgs()
// GetPublicNetworksCmd.Execute()

func TestGetPublicNetworkNotFound(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(nil, WithResponse(404, nil), nil)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.EqualError(test_framework, err, expectedMessage)
}

func TestGetPublicNetworkClientFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetPublicNetworkKeycloakFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetPublicNetworkPrinterFailure(test_framework *testing.T) {
	publicNetworkSdk := generators.Generate[networkapi.PublicNetwork]()
	publicNetworkTable := tables.PublicNetworkTableFromSdk(publicNetworkSdk)

	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(&publicNetworkSdk, WithResponse(200, WithBody(publicNetworkSdk)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(publicNetworkTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
