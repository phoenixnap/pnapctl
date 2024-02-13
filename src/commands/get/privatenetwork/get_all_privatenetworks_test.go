package privatenetwork

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllPrivateNetworksShortSuccess(test_framework *testing.T) {
	privateNetworks := testutil.GenN(5, generators.Generate[networkapi.PrivateNetwork])
	privateNetworkList := iterutils.MapInterface(privateNetworks, tables.PrivateNetworkFromSdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, nil)

	ExpectToPrintSuccess(test_framework, privateNetworkList)

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
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllPrivateNetworksPrinterFailure(test_framework *testing.T) {
	privateNetworks := testutil.GenN(5, generators.Generate[networkapi.PrivateNetwork])
	privateNetworkList := iterutils.MapInterface(privateNetworks, tables.PrivateNetworkFromSdk)

	PrepareNetworkMockClient(test_framework).
		PrivateNetworksGet("").
		Return(privateNetworks, nil)

	expectedErr := ExpectToPrintFailure(test_framework, privateNetworkList)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
