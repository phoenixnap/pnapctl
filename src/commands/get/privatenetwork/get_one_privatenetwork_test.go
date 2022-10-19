package privatenetwork

import (
	"testing"

	networkapisdk "github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetPrivateNetworkSuccess(test_framework *testing.T) {
	privateNetwork := generators.Generate[networkapisdk.PrivateNetwork]()
	var privateNetworkTable = tables.PrivateNetworkFromSdk(privateNetwork)

	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(&privateNetwork, nil)

	ExpectToPrintSuccess(test_framework, privateNetworkTable)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetPrivateNetworkClientFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetPrivateNetworkPrinterFailure(test_framework *testing.T) {
	privateNetwork := generators.Generate[networkapisdk.PrivateNetwork]()
	privateNetworkTable := tables.PrivateNetworkFromSdk(privateNetwork)

	PrepareNetworkMockClient(test_framework).
		PrivateNetworkGetById(RESOURCEID).
		Return(&privateNetwork, nil)

	expectedErr := ExpectToPrintFailure(test_framework, privateNetworkTable)

	err := GetPrivateNetworksCmd.RunE(GetPrivateNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
