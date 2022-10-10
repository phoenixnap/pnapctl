package sshkeys

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllSshKeysSuccess(test_framework *testing.T) {
	sshKeyList := testutil.GenN(2, generators.Generate[bmcapi.SshKey])

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
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllSshKeysPrinterFailure(test_framework *testing.T) {
	sshKeyList := testutil.GenN(2, generators.Generate[bmcapi.SshKey])

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
