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

func patchBgpPeerGroupSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	// What the client should receive.
	bgpPeerGroupModifyCli := generators.Generate[networkapi.BgpPeerGroupPatch]()
	bgpPeerGroupModifySdk := bgpPeerGroupModifyCli

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, bgpPeerGroupModifyCli)

	// What the server should return.
	bgpPeerGroup := generators.Generate[networkapi.BgpPeerGroup]()

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupPatchById(RESOURCEID, gomock.Eq(bgpPeerGroupModifySdk)).
		Return(&bgpPeerGroup, nil)

	// Run command
	err := PatchBgpPeerGroupCmd.RunE(PatchBgpPeerGroupCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchBgpPeerGroupSuccessYAML(test_framework *testing.T) {
	patchBgpPeerGroupSuccess(test_framework, yaml.Marshal)
}

func TestPatchBgpPeerGroupSuccessJSON(test_framework *testing.T) {
	patchBgpPeerGroupSuccess(test_framework, json.Marshal)
}

func TestPatchBgpPeerGroupFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchBgpPeerGroupCmd.RunE(PatchBgpPeerGroupCmd, []string{RESOURCEID})

	// Expected error
	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestPatchBgpPeerGroupUnmarshallingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchBgpPeerGroupCmd.RunE(PatchBgpPeerGroupCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchBgpPeerGroupClientFailure(test_framework *testing.T) {
	// What the client should receive.
	bgpPeerGroupModifyCli := generators.Generate[networkapi.BgpPeerGroupPatch]()
	bgpPeerGroupModifySdk := bgpPeerGroupModifyCli

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, bgpPeerGroupModifySdk)

	// Mocking
	PrepareNetworkMockClient(test_framework).
		BgpPeerGroupPatchById(RESOURCEID, gomock.Eq(bgpPeerGroupModifySdk)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchBgpPeerGroupCmd.RunE(PatchBgpPeerGroupCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
