package ip_blocks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetIpBlocksSuccess(test_framework *testing.T) {
	ipBlock := ipmodels.GenerateIpBlockSdk()
	tableIpBlock := tables.ToShortIpBlockTable(ipBlock)

	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(&ipBlock, WithResponse(200, WithBody(ipBlock)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tableIpBlock, "get ip-blocks").
		Return(nil)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetIpBlocksNotFound(test_framework *testing.T) {
	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(nil, WithResponse(400, nil), nil)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	expectedMessage := "Command 'get ip-blocks' has been performed, but something went wrong. Error code: 0201"
	assert.Equal(test_framework, expectedMessage, err.Error())
}

func TestGetIpBlocksClientFailure(test_framework *testing.T) {
	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(nil, nil, testutil.TestError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, "get ip-blocks", ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, expectedErr, err.Error())
}

func TestGetIpBlocksKeycloakFailure(test_framework *testing.T) {
	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetIpBlocksPrinterFailure(test_framework *testing.T) {
	ipBlock := ipmodels.GenerateIpBlockSdk()
	tableIpBlock := tables.ToShortIpBlockTable(ipBlock)

	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(&ipBlock, WithResponse(200, WithBody(tableIpBlock)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(tableIpBlock, "get ip-blocks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
