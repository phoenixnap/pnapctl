package tag

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func TestSubmitTagEditSuccessYAML(test_framework *testing.T) {
	// setup
	tag := generators.Generate[tagapi.Tag]()
	tagEdit := generators.Generate[tagapi.TagUpdate]()
	yamlmarshal, _ := yaml.Marshal(tagEdit)

	Filename = FILENAME

	//prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(tagEdit)).
		Return(&tag, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitTagEditSuccessJSON(test_framework *testing.T) {
	//setup
	tag := generators.Generate[tagapi.Tag]()
	tagEdit := generators.Generate[tagapi.TagUpdate]()
	jsonmarshal, _ := json.Marshal(tagEdit)
	Filename = FILENAME

	//prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(tagEdit)).
		Return(&tag, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

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
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())

}

func TestSubmitTagEditUnmarshallingFailure(test_framework *testing.T) {
	// setup file with incorrect data
	filecontents := []byte(`limit 45`)
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditYAMLUnmarshallingFailure(test_framework *testing.T) {
	// setup
	filecontents := []byte(`: 45`)
	yamlmarshal, _ := yaml.Marshal(filecontents)
	Filename = FILENAME

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, expectedErr, expectedErr.Error())
}

func TestSubmitTagEditFileReadingFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestSubmitTagEditClientFailure(test_framework *testing.T) {
	// setup
	tagEdit := generators.Generate[tagapi.TagUpdate]()
	yamlmarshal, _ := yaml.Marshal(tagEdit)
	Filename = FILENAME

	// prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(tagEdit)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}
