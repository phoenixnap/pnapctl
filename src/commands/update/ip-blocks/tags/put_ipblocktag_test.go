package ip_blocks

import (
	"encoding/json"
	"testing"

	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
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
	yamlmarshal, _ := yaml.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(yamlmarshal, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestPutIpBlockTagSuccessJSON(test_framework *testing.T) {
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(&ipBlock, nil)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestIpBlockPutTagFileNotFoundFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIValidationError{Message: "The file '" + FILENAME + "' does not exist."})

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected command
	expectedErr := ctlerrors.FileNotExistError(FILENAME)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestIpBlockPutTagUnmarshallingFailure(test_framework *testing.T) {
	// Invalid contents of the file
	filecontents := []byte(`error error`)

	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(filecontents, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.UnmarshallingInFileProcessor, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestIpBlockPutTagFileReadingFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(nil, ctlerrors.CLIError{
			Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: 0503",
		})

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.CreateCLIError(ctlerrors.FileReading, err)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestIpBlockPutTagClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockPutTagCli := testutil.GenN(3, generators.Generate[ipapi.TagAssignmentRequest])

	// Assumed contents of the file.
	jsonmarshal, _ := json.Marshal(ipBlockPutTagCli)

	Filename = FILENAME

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksIpBlockIdTagsPut(RESOURCEID, gomock.Eq(ipBlockPutTagCli)).
		Return(nil, testutil.TestError)

	PrepareMockFileProcessor(test_framework).
		ReadFile(FILENAME).
		Return(jsonmarshal, nil)

	// Run command
	err := PutIpBlockTagCmd.RunE(PutIpBlockTagCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
