package sshkeys

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/tables"
	"phoenixnap.com/pnap-cli/tests/generators"
	. "phoenixnap.com/pnap-cli/tests/mockhelp"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

func TestGetSshKeyByIdFullSuccess(test_framework *testing.T) {
	Full = true
	sshKey := generators.GenerateSshKey()
	sshKeyTable := tables.ToSshKeyTableFull(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(sshKey, WithResponse(200, WithBody(sshKey)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable, "get ssh-keys").
		Return(nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetSshKeyByIdSuccess(test_framework *testing.T) {
	Full = false
	sshKey := generators.GenerateSshKey()
	sshKeyTable := tables.ToSshKeyTable(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(sshKey, WithResponse(200, WithBody(sshKey)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable, "get ssh-keys").
		Return(nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetSshKeyByIdNotFound(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(bmcapisdk.SshKey{}, WithResponse(400, nil), nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get ssh-keys' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetSshKeyByIdClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(bmcapisdk.SshKey{}, nil, testutil.TestError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get ssh-keys", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetSshKeyByIdKeycloakFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(bmcapisdk.SshKey{}, nil, testutil.TestKeycloakError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetSshKeyByIdPrinterFailure(test_framework *testing.T) {
	Full = false
	sshKey := generators.GenerateSshKey()
	sshKeyTable := tables.ToSshKeyTable(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(sshKey, WithResponse(200, WithBody(sshKey)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable, "get ssh-keys").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
