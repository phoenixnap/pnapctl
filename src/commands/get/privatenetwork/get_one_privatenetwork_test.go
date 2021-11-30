package privatenetwork

import (
	"errors"
	"testing"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/tests/generators"
	. "phoenixnap.com/pnapctl/tests/mockhelp"
	"phoenixnap.com/pnapctl/tests/testutil"
)

func TestGetPrivateNetworkSuccess(test_framework *testing.T) {

	privateNetwork := generators.GeneratePrivateNetwork()
	var privateNetworkTable = tables.PrivateNetworkFromSdk(privateNetwork)

	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkTable, "get private-network").
		Return(nil)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetPrivateNetworkNotFound(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(networksdk.PrivateNetwork{}, WithResponse(400, nil), nil)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get private-network' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetPrivateNetworkClientFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get private-network", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetPrivateNetworkKeycloakFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(networksdk.PrivateNetwork{}, nil, testutil.TestKeycloakError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetPrivateNetworkPrinterFailure(test_framework *testing.T) {
	privateNetwork := generators.GeneratePrivateNetwork()
	privateNetworkTable := tables.PrivateNetworkFromSdk(privateNetwork)

	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(privateNetwork, WithResponse(200, WithBody(privateNetwork)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkTable, "get private-network").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
