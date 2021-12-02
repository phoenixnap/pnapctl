package privatenetwork

import (
	"errors"
	"testing"

	networksdk "github.com/phoenixnap/go-sdk-bmc/networkapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/tests/mockhelp"
	"phoenixnap.com/pnapctl/tests/testutil"
)

func TestGetAllPrivateNetworksShortSuccess(test_framework *testing.T) {
	privateNetworks := networkmodels.GeneratePrivateNetworks(5)

	var privateNetworkList []interface{}

	for _, x := range privateNetworks {
		privateNetworkList = append(privateNetworkList, tables.PrivateNetworkFromSdk(x))
	}

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, WithResponse(200, WithBody(privateNetworks)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkList, "get private-network").
		Return(nil)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllPrivateNetworksClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return([]networksdk.PrivateNetwork{}, WithResponse(200, nil), testutil.TestError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get servers", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllPrivateNetworksKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return([]networksdk.PrivateNetwork{}, nil, testutil.TestKeycloakError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllPrivateNetworksPrinterFailure(test_framework *testing.T) {
	privateNetworks := networkmodels.GeneratePrivateNetworks(5)

	var privateNetworkList []interface{}

	for _, x := range privateNetworks {
		privateNetworkList = append(privateNetworkList, tables.PrivateNetworkFromSdk(x))
	}

	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, WithResponse(200, WithBody(privateNetworks)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkList, "get private-network").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
