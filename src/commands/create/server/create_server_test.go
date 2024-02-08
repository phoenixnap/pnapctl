package server

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"

	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
)

func getQueryParams() bool {
	return force
}

func createServerSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	serverCreate := generators.Generate[bmcapisdk.ServerCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverCreate)

	// What the server should return.
	createdServer := generators.Generate[bmcapisdk.Server]()

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate), force).
		Return(&createdServer, nil)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerSuccessYAML(test_framework *testing.T) {
	createServerSuccess(test_framework, yaml.Marshal)
}

func TestCreateServerSuccessJSON(test_framework *testing.T) {
	createServerSuccess(test_framework, json.Marshal)
}

func TestCreateServerFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateServerClientFailure(test_framework *testing.T) {
	// Setup
	serverCreate := generators.Generate[bmcapisdk.ServerCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, serverCreate)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersPost(gomock.Eq(serverCreate), force).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateServerCmd.RunE(CreateServerCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
