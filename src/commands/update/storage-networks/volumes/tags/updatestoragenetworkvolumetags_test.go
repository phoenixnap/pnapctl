package tags

import (
	"encoding/json"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func updateStorageNetworkVolumeTagsSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	tagAssignmentRequest := generators.Generate[[]networkstorageapi.TagAssignmentRequest]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, tagAssignmentRequest)

	// What the server should return.
	volume := generators.Generate[networkstorageapi.Volume]()

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageVolumePutTags(RESOURCEID, RESOURCEID, tagAssignmentRequest).
		Return(&volume, nil)

	// Run command
	err := UpdateStorageNetworkVolumeTagsCmd.RunE(UpdateStorageNetworkVolumeTagsCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestUpdateStorageNetworkVolumeTagsSuccessYAML(test_framework *testing.T) {
	updateStorageNetworkVolumeTagsSuccess(test_framework, yaml.Marshal)
}

func TestUpdateStorageNetworkVolumeTagsSuccessJSON(test_framework *testing.T) {
	updateStorageNetworkVolumeTagsSuccess(test_framework, json.Marshal)
}

func TestUpdateStorageNetworkVolumeTagsFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := UpdateStorageNetworkVolumeTagsCmd.RunE(UpdateStorageNetworkVolumeTagsCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestUpdateStorageNetworkVolumeTagsUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := UpdateStorageNetworkVolumeTagsCmd.RunE(UpdateStorageNetworkVolumeTagsCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestUpdateStorageNetworkVolumeTagsClientFailure(test_framework *testing.T) {
	// What the client should receive.
	tagAssignmentRequest := generators.Generate[[]networkstorageapi.TagAssignmentRequest]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, tagAssignmentRequest)

	// Mocking
	PrepareNetworkStorageApiMockClient(test_framework).
		NetworkStorageVolumePutTags(RESOURCEID, RESOURCEID, tagAssignmentRequest).
		Return(nil, testutil.TestError)

	// Run command
	err := UpdateStorageNetworkVolumeTagsCmd.RunE(UpdateStorageNetworkVolumeTagsCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
