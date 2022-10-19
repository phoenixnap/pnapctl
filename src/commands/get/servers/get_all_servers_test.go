package servers

import (
	"errors"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllServersShortSuccess(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	shortServers := iterutils.Map(serverlist, tables.ToShortServerTable)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServers).
		Return(nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	longServers := iterutils.Map(serverlist, tables.ToLongServerTable)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServers).
		Return(nil)

	// to display full output
	Full = true

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestFilteredServersLongSuccess(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	longServers := iterutils.Map(serverlist, tables.ToLongServerTable)

	// to display full output
	Full = true
	tags = []string{"tag1.value1", "tag2.value2"}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(longServers).
		Return(nil)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersClientFailure(test_framework *testing.T) {
	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(nil, testutil.TestError)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetAllServersPrinterFailure(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	shortServers := iterutils.Map(serverlist, tables.ToShortServerTable)

	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(shortServers).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	Full = false

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
