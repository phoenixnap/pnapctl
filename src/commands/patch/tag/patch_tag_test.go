package tag

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func submitTagEditSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// setup
	tag := generators.Generate[tagapi.Tag]()
	tagEdit := generators.Generate[tagapi.TagUpdate]()
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, tagEdit)

	//prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(tagEdit)).
		Return(&tag, nil)

	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.NoError(test_framework, err)
}

func TestSubmitTagEditSuccessYAML(test_framework *testing.T) {
	submitTagEditSuccess(test_framework, yaml.Marshal)
}

func TestSubmitTagEditSuccessJSON(test_framework *testing.T) {
	submitTagEditSuccess(test_framework, json.Marshal)
}

func TestSubmitTagEditFileProcessorFailure(test_framework *testing.T) {
	// setup
	Filename = FILENAME

	// prepare mocks
	expectedErr := ExpectFromFileFailure(test_framework)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestSubmitTagEditUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// prepare mocks
	ExpectFromFileUnmarshalFailure(test_framework)

	// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestSubmitTagEditClientFailure(test_framework *testing.T) {
	// setup
	tagEdit := generators.Generate[tagapi.TagUpdate]()
	ExpectFromFileSuccess(test_framework, yaml.Marshal, tagEdit)
	Filename = FILENAME

	// prepare mocks
	PrepareTagMockClient(test_framework).
		TagPatch(RESOURCEID, gomock.Eq(tagEdit)).
		Return(nil, testutil.TestError)

		// execute
	err := PatchTagCmd.RunE(PatchTagCmd, []string{RESOURCEID})

	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
