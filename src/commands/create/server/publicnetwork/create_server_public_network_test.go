package publicnetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func createReservationSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverPublicNetwork)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(serverPublicNetwork)).
		Return(&serverPublicNetwork, nil)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPublicNetworkSuccessYAML(test_framework *testing.T) {
	createReservationSuccess(test_framework, yaml.Marshal)
}

func TestCreateServerPublicNetworkSuccessJSON(test_framework *testing.T) {
	createReservationSuccess(test_framework, json.Marshal)
}

func TestCreateServerPublicNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerPublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateServerPublicNetworkClientFailure(test_framework *testing.T) {
	// Setup
	serverPublicNetwork := generators.Generate[bmcapisdk.ServerPublicNetwork]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPublicNetwork)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPublicNetworkPost(RESOURCEID, gomock.Eq(serverPublicNetwork)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateServerPublicNetworkCmd.RunE(CreateServerPublicNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
