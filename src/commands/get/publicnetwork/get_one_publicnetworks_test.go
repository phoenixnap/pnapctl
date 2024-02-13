package publicnetwork

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetPublicNetworkSuccess(test_framework *testing.T) {
	publicNetworkSdk := generators.Generate[networkapi.PublicNetwork]()
	publicNetworkTable := tables.PublicNetworkTableFromSdk(publicNetworkSdk)

	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(&publicNetworkSdk, nil)

	ExpectToPrintSuccess(test_framework, publicNetworkTable)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetPublicNetworkClientFailure(test_framework *testing.T) {
	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetPublicNetworkPrinterFailure(test_framework *testing.T) {
	publicNetworkSdk := generators.Generate[networkapi.PublicNetwork]()
	publicNetworkTable := tables.PublicNetworkTableFromSdk(publicNetworkSdk)

	PrepareNetworkMockClient(test_framework).
		PublicNetworkGetById(RESOURCEID).
		Return(&publicNetworkSdk, nil)

	expectedErr := ExpectToPrintFailure(test_framework, publicNetworkTable)

	err := GetPublicNetworksCmd.RunE(GetPublicNetworksCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
