package privatenetwork

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllPrivateNetworksShortSuccess(test_framework *testing.T) {
	privateNetworks := testutil.GenN(5, generators.Generate[networkapi.PrivateNetwork])

	var privateNetworkList []interface{}

	for _, x := range privateNetworks {
		privateNetworkList = append(privateNetworkList, tables.PrivateNetworkFromSdk(x))
	}

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkList).
		Return(nil)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllPrivateNetworksClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(nil, testutil.TestError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetAllPrivateNetworksPrinterFailure(test_framework *testing.T) {
	privateNetworks := testutil.GenN(5, generators.Generate[networkapi.PrivateNetwork])

	var privateNetworkList []interface{}

	for _, x := range privateNetworks {
		privateNetworkList = append(privateNetworkList, tables.PrivateNetworkFromSdk(x))
	}

	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(privateNetworkList).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
