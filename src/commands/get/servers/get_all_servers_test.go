package servers

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
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
	shortServers := iterutils.MapInterface(serverlist, tables.ToShortServerTable)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	ExpectToPrintSuccess(test_framework, shortServers)

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllServersLongSuccess(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	longServers := iterutils.MapInterface(serverlist, tables.ToLongServerTable)

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	ExpectToPrintSuccess(test_framework, longServers)

	// to display full output
	Full = true

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestFilteredServersLongSuccess(test_framework *testing.T) {
	serverlist := testutil.GenN(5, generators.Generate[bmcapisdk.Server])
	longServers := iterutils.MapInterface(serverlist, tables.ToLongServerTable)

	// to display full output
	Full = true
	tags = []string{"tag1.value1", "tag2.value2"}

	// Mocking
	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	ExpectToPrintSuccess(test_framework, longServers)

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
	shortServers := iterutils.MapInterface(serverlist, tables.ToShortServerTable)

	PrepareBmcApiMockClient(test_framework).
		ServersGet(tags).
		Return(serverlist, nil)

	expectedErr := ExpectToPrintFailure(test_framework, shortServers)

	Full = false

	err := GetServersCmd.RunE(GetServersCmd, []string{})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
