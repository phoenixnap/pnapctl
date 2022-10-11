package ipblocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

const deleteResult = "The specified IP block is being removed from the server."

func TestDeleteServerIpBlockSuccessYAML(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return(deleteResult, WithResponse(202, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerIpBlockSuccessJSON(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return(deleteResult, WithResponse(202, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerIpBlockNotFound(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", WithResponse(404, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerIpBlockError(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", WithResponse(500, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	expectedMessage := "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerIpBlockSdkentFailure(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", nil, testutil.TestError)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerIpBlockKeycloakFailure(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", nil, testutil.TestKeycloakError)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
