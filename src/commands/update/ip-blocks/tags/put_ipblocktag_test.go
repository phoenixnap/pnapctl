package ip_blocks

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

func TestPutIpBlockTagSuccessYAML(test_framework *testing.T) {
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, ipBlockPutTagCli)

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPutIpBlockTagSuccessJSON(test_framework *testing.T) {
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, ipBlockPutTagCli)

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestIpBlockPutTagFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestIpBlockPutTagUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestIpBlockPutTagClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, ipBlockPutTagCli)

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, testutil.TestError)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
