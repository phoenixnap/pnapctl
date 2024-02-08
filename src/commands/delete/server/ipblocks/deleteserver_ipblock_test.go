package ipblocks

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"sigs.k8s.io/yaml"
)

const deleteResult = "The specified IP block is being removed from the server."

func createReservationSuccess(test_framework *testing.T, marshaller func(interface{}) ([]byte, error)) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, marshaller, relinquishIpBlock)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return(deleteResult, nil)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestDeleteServerIpBlockSuccessYAML(test_framework *testing.T) {
	createReservationSuccess(test_framework, yaml.Marshal)
}

func TestDeleteServerIpBlockSuccessJSON(test_framework *testing.T) {
	createReservationSuccess(test_framework, json.Marshal)
}

func TestDeleteServerIpBlockSuccessUnmarshallingError(test_framework *testing.T) {
	Filename = FILENAME

	// Mocking
	expectedErr := ExpectFromFileFailure(test_framework)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Assertions
	assert.Equal(test_framework, expectedErr, err)
}

func TestDeleteServerIpBlockClientFailure(test_framework *testing.T) {
	relinquishIpBlock := generators.Generate[bmcapisdk.RelinquishIpBlock]()

	// Assumed contents of the file.
	Filename = FILENAME
	ExpectFromFileSuccess(test_framework, yaml.Marshal, relinquishIpBlock)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServerIpBlockDelete(RESOURCEID, RESOURCEID, gomock.Eq(relinquishIpBlock)).
		Return("", testutil.TestError)

	// Run command
	err := DeleteServerIpBlockCmd.RunE(DeleteServerIpBlockCmd, []string{RESOURCEID, RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(testutil.TestError, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
