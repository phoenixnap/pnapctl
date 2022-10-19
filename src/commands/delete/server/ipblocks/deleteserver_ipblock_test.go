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

const deleteResult = "The specified IP block is being removed from the server."

func TestDeleteServerIpBlockSuccessYAML(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return(deleteResult, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

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
		Return(deleteResult, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerIpBlockSuccessUnmarshallingError(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, testutil.TestError)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestDeleteServerIpBlockClientFailure(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	yamlmarshal, _ := yaml.Marshal(relinquishIpBlock)

	Filename = FILENAME

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
