package ipblock

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createPublicNetworkIpBlockSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	ipBlockCreate := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, ipBlockCreate)

	// What the server should return.
	createdIpBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(ipBlockCreate)).
		Return(&createdIpBlock, nil)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkIpBlockSuccessYAML(test_framework *testing.T) {
	createPublicNetworkIpBlockSuccess(test_framework, yaml.Marshal)
}

func TestCreatePublicNetworkIpBlockSuccessJSON(test_framework *testing.T) {
	createPublicNetworkIpBlockSuccess(test_framework, json.Marshal)
}

func TestCreatePublicNetworkIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkIpBlockUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, ipBlockCreate)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(ipBlockCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
