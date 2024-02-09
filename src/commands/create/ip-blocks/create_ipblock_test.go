package ip_blocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

func createIpBlockSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	ipBlockCreate := generators.Generate[ipapi.IpBlockCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, ipBlockCreate)

	// What the server should return.
	ipBlock := generators.Generate[ipapi.IpBlock]()

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(&ipBlock, nil)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestCreateIpBlockSuccessYAML(test_framework *testing.T) {
	createIpBlockSuccess(test_framework, yaml.Marshal)
}

func TestCreateIpBlockSuccessJSON(test_framework *testing.T) {
	createIpBlockSuccess(test_framework, json.Marshal)
}

func TestCreateIpBlockFileProcessorFailure(test_framework *testing.T) {
	// Setup
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())

}

func TestCreateIpBlockUnmarshallingFailure(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	ExpectFromFileUnmarshalFailure(test_framework)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInFileProcessor)
}

func TestCreateIpBlockClientFailure(test_framework *testing.T) {
	// Setup
	ipBlockCreate := generators.Generate[ipapi.IpBlockCreate]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, json.Marshal, ipBlockCreate)

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlockPost(gomock.Eq(ipBlockCreate)).
		Return(nil, testutil.TestError)

	// Run command
	err := CreateIpBlockCmd.RunE(CreateIpBlockCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
