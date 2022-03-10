package ipblocks

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"testing"
)

const deleteResult = "The specified IP block is being removed from the server."

func TestDeleteServerIpBlockSuccessYAML(test_framework *testing.T) {
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
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
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
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
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
		Return("", WithResponse(404, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	expectedMessage := "Command 'delete server-ip-block' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())

}

func TestDeleteServerIpBlockError(test_framework *testing.T) {
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
		Return("", WithResponse(500, nil), nil)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	expectedMessage := "Command 'delete server-ip-block' has been performed, but something went wrong. Error code: 0201"

	// Assertions
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestDeleteServerIpBlockClientFailure(test_framework *testing.T) {
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
		Return("", nil, testutil.TestError)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "delete server-ip-block", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestDeleteServerIpBlockKeycloakFailure(test_framework *testing.T) {
	relinquishIpBlock := servermodels.GenerateRelinquishIpBlockCli()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(*relinquishIpBlock.ToSdk())).
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
