package ipblocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createServerIpBlockSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	serverIpBlockSdk := generators.Generate[bmcapisdk.ServerIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, serverIpBlockSdk)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(&serverIpBlockSdk, nil)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateServerIpBlockSuccessYAML(test_framework *testing.T) {
	createServerIpBlockSuccess(test_framework, yaml.Marshal)
}

func TestCreateServerIpBlockSuccessJSON(test_framework *testing.T) {
	createServerIpBlockSuccess(test_framework, json.Marshal)
}

func TestCreateServerIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateServerIpBlockUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateServerIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	serverIpBlockSdk := generators.Generate[bmcapisdk.ServerIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, serverIpBlockSdk)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockPost(RESOURCEID, gomock.Eq(serverIpBlockSdk)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateServerIpBlockCmd.RunE(CreateServerIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
