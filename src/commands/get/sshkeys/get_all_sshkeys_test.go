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

func TestGetAllSshKeysSuccess(test_framework *testing.T) {
	sshKeyList := generators.GenerateSshKeys(2)

	var sshKeyTables []interface{}

	for _, sshKey := range sshKeyList {
		sshKeyTables = append(sshKeyTables, tables.ToSshKeyTable(sshKey))
	}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(sshKeyList, WithResponse(200, WithBody(sshKeyList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTables, "get ssh-keys").
		Return(nil)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllSshKeysKeycloakFailure(test_framework *testing.T) {
	sshKey := []bmcapisdk.SshKey{generators.GenerateSshKey()}
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(sshKey, nil, testutil.TestKeycloakError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllSshKeysPrinterFailure(test_framework *testing.T) {
	sshKeyList := generators.GenerateSshKeys(2)

	var sshKeyTables []interface{}

	for _, sshKey := range sshKeyList {
		sshKeyTables = append(sshKeyTables, tables.ToSshKeyTable(sshKey))
	}

	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(sshKeyList, WithResponse(200, WithBody(sshKeyList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(sshKeyTables, "get ssh-keys").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
