package quotas

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestSubmitQuotaEditRequestSuccessYAML(test_framework *testing.T) {
	// setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(quotaEditLimitRequest)

	Filename = FILENAME

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(WithResponse(202, WithBody(nil)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestSuccessJSON(test_framework *testing.T) {
	//setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	jsonmarshal, _ := json.Marshal(quotaEditLimitRequest)
	Filename = FILENAME

	//prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(WithResponse(202, WithBody(nil)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(jsonmarshal, nil).
		Times(1)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitQuotaEditRequestFileNotFoundFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{})

	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestSubmitQuotaEditRequestUnmarshallingFailure(test_framework *testing.T) {
	// setup file with incorrect data
	filecontents := []byte(`limit 45`)
	Filename = FILENAME

	// prepare mocks
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(filecontents, nil).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "request-edit quota", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitQuotaEditRequestYAMLUnmarshallingFailure(test_framework *testing.T) {
	// setup
	filecontents := []byte(`Limit: 45`)
	yamlmarshal, _ := yaml.Marshal(filecontents)
	Filename = FILENAME

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "request-edit quota", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, expectedErr.Error())
}

func TestSubmitQuotaEditFileReadingFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'request-edit quota' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "request-edit quota", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitQuotaEditBackendErrorFailure(test_framework *testing.T) {
	// setup
	quotaEditLimitRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(quotaEditLimitRequest)
	Filename = FILENAME

	// prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(quotaEditLimitRequest)).
		Return(WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitQuotaEditClientFailure(test_framework *testing.T) {
	// setup
	editQuotaRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(editQuotaRequest)
	Filename = FILENAME

	// prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(editQuotaRequest)).
		Return(nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "request-edit quota", ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitQuotaEditKeycloakFailure(test_framework *testing.T) {
	// setup
	editQuotaRequest := generators.Generate[bmcapi.QuotaEditLimitRequest]()
	yamlmarshal, _ := yaml.Marshal(editQuotaRequest)
	Filename = FILENAME

	// prepare mocks
	PrepareBmcApiMockClient(test_framework).
		QuotaEditById(RESOURCEID, gomock.Eq(editQuotaRequest)).
		Return(nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME, commandName).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := RequestEditQuotaCmd.RunE(RequestEditQuotaCmd, []string{RESOURCEID})

	// assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
