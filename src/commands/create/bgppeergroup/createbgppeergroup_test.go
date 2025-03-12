package bgppeergroup

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v4"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createBgpPeerGroupSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	bgpPeerGroupCreate := generators.Generate[networkapi.BgpPeerGroupCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, bgpPeerGroupCreate)

	// What the server should return.
	createdBgpPeerGroup := generators.Generate[networkapi.BgpPeerGroup]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupsPost(gomock.Eq(bgpPeerGroupCreate)).
		Return(&createdBgpPeerGroup, nil)

	// Run command
	err := CreateBgpPeerGroupCmd.RunE(CreateBgpPeerGroupCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateBgpPeerGroupSuccessYAML(test_framework *testing.T) {
	createBgpPeerGroupSuccess(test_framework, yaml.Marshal)
}

func TestCreateBgpPeerGroupSuccessJSON(test_framework *testing.T) {
	createBgpPeerGroupSuccess(test_framework, json.Marshal)
}

func TestCreateBgpPeerGroupFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateBgpPeerGroupCmd.RunE(CreateBgpPeerGroupCmd, []string{})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestCreateBgpPeerGroupUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateBgpPeerGroupCmd.RunE(CreateBgpPeerGroupCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateBgpPeerGroupClientFailure(test_framework *testing.T) {
	// What the client should receive.
	bgpPeerGroupCreate := generators.Generate[networkapi.BgpPeerGroupCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, bgpPeerGroupCreate)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupsPost(gomock.Eq(bgpPeerGroupCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateBgpPeerGroupCmd.RunE(CreateBgpPeerGroupCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
