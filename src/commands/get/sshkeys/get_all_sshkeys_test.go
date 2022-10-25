package sshkeys

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllSshKeysSuccess(test_framework *testing.T) {
	sshKeyList := testutil.GenN(2, generators.Generate[bmcapi.SshKey])
	sshKeyTables := iterutils.MapInterface(sshKeyList, tables.ToSshKeyTable)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(sshKeyList, nil)

	ExpectToPrintSuccess(test_framework, sshKeyTables)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllSshKeysClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(nil, testutil.TestError)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllSshKeysPrinterFailure(test_framework *testing.T) {
	sshKeyList := testutil.GenN(2, generators.Generate[bmcapi.SshKey])
	sshKeyTables := iterutils.MapInterface(sshKeyList, tables.ToSshKeyTable)

	PrepareBmcApiMockClient(test_framework).
		SshKeysGet().
		Return(sshKeyList, nil)

	expectedErr := ExpectToPrintFailure(test_framework, sshKeyTables)

	err := GetSshKeysCmd.RunE(GetSshKeysCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
