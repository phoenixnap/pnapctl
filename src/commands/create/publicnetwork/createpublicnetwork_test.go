package publicnetwork

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

func TestCreatePublicNetworkSuccessYAML(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkSuccessJSON(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// What the server should return.
	createdPublicNetwork := generators.Generate[networkapi.PublicNetwork]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(&createdPublicNetwork, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreatePublicNetworkFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileFailure(test_framework)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := testutil.TestError

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	filecontents := []byte(`invalid`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreatePublicNetworkFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreatePublicNetworkClientFailure(test_framework *testing.T) {
	// What the client should receive.
	publicNetworkCreate := generators.Generate[networkapi.PublicNetworkCreate]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(publicNetworkCreate)

	Filename = FILENAME

	// Mocking
	PrepareNetworkMockClient(test_framework).
		PublicNetworksPost(gomock.Eq(publicNetworkCreate)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := CreatePublicNetworkCmd.RunE(CreatePublicNetworkCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
