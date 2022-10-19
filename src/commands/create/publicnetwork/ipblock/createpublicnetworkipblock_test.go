package ipblock

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestCreatePublicNetworkIpBlockSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// What the server should return.
	createdIpBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(ipBlockCreate)).
		Return(&createdIpBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkIpBlockSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockCreate)

	Filename = FILENAME

	// What the server should return.
	createdIpBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(ipBlockCreate)).
		Return(&createdIpBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkIpBlockUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	filecontents := []byte(`invalid`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePublicNetworkIpBlockFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkIpBlockClientFailure(test_framework *testing.T) {
	// What the client should receive.
	ipBlockCreate := generators.Generate[networkapi.PublicNetworkIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(ipBlockCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworkIpBlockPost(RESOURCEID, gomock.Eq(ipBlockCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePublicNetworkIpBlockCmd.RunE(CreatePublicNetworkIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
