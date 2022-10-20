package privatenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func createServerPrivateNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverPrivateNetwork := generators.Generate[bmcapisdk.ServerPrivateNetwork]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, serverPrivateNetwork)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(serverPrivateNetwork)).
		Return(&serverPrivateNetwork, nil)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerPrivateNetworkSuccessYAML(test_framework *testing.T) {
	createServerPrivateNetworkSuccess(test_framework, yaml.Marshal)
}

func TestCreateServerPrivateNetworkSuccessJSON(test_framework *testing.T) {
	createServerPrivateNetworkSuccess(test_framework, json.Marshal)
}

func TestCreateServerPrivateNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerPrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateServerPrivateNetworkClientFailure(test_framework *testing.T) {
	// Setup
	serverPrivateNetwork := generators.Generate[bmcapisdk.ServerPrivateNetwork]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverPrivateNetwork)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerPrivateNetworkPost(RESOURCEID, gomock.Eq(serverPrivateNetwork)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateServerPrivateNetworkCmd.RunE(CreateServerPrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
