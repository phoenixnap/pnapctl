package tag

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestSubmitTagEditSuccessYAML(test_framework *testing.T) {
	// setup
	tag := *tagmodels.GenerateTagSdk()
	tagEdit := *tagmodels.GenerateTagUpdateCli()
	yamlmarshal, _ := yaml.Marshal(tagEdit)

	Filename = FILENAME

	//prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(*tagEdit.ToSdk())).
		Return(&tag, WithResponse(200, WithBody(nil)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitTagEditSuccessJSON(test_framework *testing.T) {
	//setup
	tag := *tagmodels.GenerateTagSdk()
	tagEdit := *tagmodels.GenerateTagUpdateCli()
	jsonmarshal, _ := json.Marshal(tagEdit)
	Filename = FILENAME

	//prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(*tagEdit.ToSdk())).
		Return(&tag, WithResponse(200, WithBody(nil)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(jsonmarshal, nil).
		Times(1)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitTagEditFileNotFoundFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."}).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{})

	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestSubmitTagEditUnmarshallingFailure(test_framework *testing.T) {
	// setup file with incorrect data
	filecontents := []byte(`limit 45`)
	Filename = FILENAME

	// prepare mocks
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(filecontents, nil).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch tag", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditYAMLUnmarshallingFailure(test_framework *testing.T) {
	// setup
	filecontents := []byte(`: 45`)
	yamlmarshal, _ := yaml.Marshal(filecontents)
	Filename = FILENAME

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, "patch tag", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, expectedErr.Error())
}

func TestSubmitTagEditFileReadingFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command 'patch tag' has been performed, but something went wrong. Error code: 0503",
		}).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, "patch tag", err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditBackendErrorFailure(test_framework *testing.T) {
	// setup
	tagEdit := *tagmodels.GenerateTagUpdateCli()
	yamlmarshal, _ := yaml.Marshal(&tagEdit)
	Filename = FILENAME

	// prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(*tagEdit.ToSdk())).
		Return(nil, WithResponse(500, WithBody(testutil.GenericBMCError)), nil).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := errors.New(testutil.GenericBMCError.Message)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditClientFailure(test_framework *testing.T) {
	// setup
	tagEdit := *tagmodels.GenerateTagUpdateCli()
	yamlmarshal, _ := yaml.Marshal(tagEdit)
	Filename = FILENAME

	// prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(*tagEdit.ToSdk())).
		Return(nil, nil, testutil.TestError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, "patch tag", ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditKeycloakFailure(test_framework *testing.T) {
	// setup
	tagEdit := *tagmodels.GenerateTagUpdateCli()
	yamlmarshal, _ := yaml.Marshal(tagEdit)
	Filename = FILENAME

	// prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(*tagEdit.ToSdk())).
		Return(nil, nil, testutil.TestKeycloakError).
		Times(1)

	mockFileProcessor := PrepareMockFileProcessor(test_framework)

	mockFileProcessor.
		ReadFile(FILENAME).
		Return(yamlmarshal, nil).
		Times(1)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}
