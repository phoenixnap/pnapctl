package ipblock

import (
	"encoding/json"
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/stretchr/testify/assert"

	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"sigs.k8s.io/yaml"
)

func TestPatchIpBlockSuccessYAML(test_framework *testing.T) {
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, ipBlockPatchCli)

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(&ipBlock, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchIpBlockSuccessJSON(test_framework *testing.T) {
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, ipBlockPatchCli)

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(&ipBlock, nil)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPatchIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestPatchIpBlockUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestPatchIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockPatchCli := generators.Generate[ipapi.IpBlockPatch]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, ipBlockPatchCli)

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdPatch(RESOURCEID, gomock.Eq(ipBlockPatchCli)).
		Return(nil, testutil.TestError)

	// Run command
	err := PatchIpBlockCmd.RunE(PatchIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
