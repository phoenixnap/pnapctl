package sshkeys

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetSshKeyByIdFullSuccess(test_framework *testing.T) {
	Full = true
	sshKey := generators.Generate[bmcapisdk.SshKey]()
	sshKeyTable := tables.ToSshKeyTableFull(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(&sshKey, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable).
		Return(nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetSshKeyByIdSuccess(test_framework *testing.T) {
	Full = false
	sshKey := generators.Generate[bmcapisdk.SshKey]()
	sshKeyTable := tables.ToSshKeyTable(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(&sshKey, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable).
		Return(nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetSshKeyByIdClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetSshKeyByIdPrinterFailure(test_framework *testing.T) {
	Full = false
	sshKey := generators.Generate[bmcapisdk.SshKey]()
	sshKeyTable := tables.ToSshKeyTable(sshKey)

	PrepareBmcApiMockClient(test_framework).
		SshKeyGetById(RESOURCEID).
		Return(&sshKey, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTable).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
