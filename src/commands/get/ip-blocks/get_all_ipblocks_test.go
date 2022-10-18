package ip_blocks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllIpBlocksSuccess(test_framework *testing.T) {
	ipBlockList := generators.GenerateIpBlockSdkList(2)

	var IpBlockTables []interface{}

	for _, ipBlock := range ipBlockList {
		IpBlockTables = append(IpBlockTables, tables.ToShortIpBlockTable(ipBlock))
	}

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(ipBlockList, WithResponse(200, WithBody(ipBlockList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(IpBlockTables, "get ip-blocks").
		Return(nil)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllIpBlocksKeycloakFailure(test_framework *testing.T) {
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(nil, nil, testutil.TestKeycloakError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllIpBlocksPrinterFailure(test_framework *testing.T) {
	ipBlockList := generators.GenerateIpBlockSdkList(2)

	var ipBlockTables []interface{}

	for _, ipBlock := range ipBlockList {
		ipBlockTables = append(ipBlockTables, tables.ToShortIpBlockTable(ipBlock))
	}

	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(ipBlockList, WithResponse(200, WithBody(ipBlockList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(ipBlockTables, "get ip-blocks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
