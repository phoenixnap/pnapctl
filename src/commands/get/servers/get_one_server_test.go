package servers

import (
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetServerShortSuccess(test_framework *testing.T) {
	server := generators.Generate[bmcapisdk.Server]()
	var shortServer = tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(&server, nil)

	ExpectToPrintSuccess(test_framework, shortServer)

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerLongSuccess(test_framework *testing.T) {
	server := generators.Generate[bmcapisdk.Server]()
	var longServer = tables.ToLongServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(&server, nil)

	ExpectToPrintSuccess(test_framework, longServer)

	Full = true
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetServerClientFailure(test_framework *testing.T) {
	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetServerPrinterFailure(test_framework *testing.T) {
	server := generators.Generate[bmcapisdk.Server]()
	shortServer := tables.ToShortServerTable(server)

	PrepareBmcApiMockClient(test_framework).
		ServerGetById(RESOURCEID).
		Return(&server, nil)

	expectedErr := ExpectToPrintFailure(test_framework, shortServer)

	Full = false
	err := GetServersCmd.RunE(GetServersCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
