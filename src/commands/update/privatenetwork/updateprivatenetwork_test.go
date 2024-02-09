package privatenetwork

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func updatePrivateNetworkSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, privateNetworkUpdate)

	// What the server should return.
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(&privateNetwork, nil)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdatePrivateNetworkSuccessYAML(test_framework *testing.T) {
	updatePrivateNetworkSuccess(test_framework, yaml.Marshal)
}

func TestUpdatePrivateNetworkSuccessJSON(test_framework *testing.T) {
	updatePrivateNetworkSuccess(test_framework, json.Marshal)
}

func TestUpdatePrivateNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestUpdatePrivateNetworkUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestUpdatePrivateNetworkClientFailure(test_framework *testing.T) {
	// Setup
	privateNetworkUpdate := generators.Generate[networkapi.PrivateNetworkModify]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, privateNetworkUpdate)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PrivateNetworkPut(RESOURCEID, gomock.Eq(privateNetworkUpdate)).
		Return(nil, testutil.TestError)

	// Run command
	err := UpdatePrivateNetworkCmd.RunE(UpdatePrivateNetworkCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
