package ip_blocks

import (
	"errors"
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/ipmodels"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
	"testing"
)

func TestGetAllIpBlocksSuccess(test_framework *testing.T) {
	ipBlockList := ipmodels.GenerateIpBlockSdkList(2)

	var IpBlockTables []interface{}

	for _, ipBlock := range ipBlockList {
		IpBlockTables = append(IpBlockTables, tables.ToIpBlockTable(ipBlock))
	}

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksGet().
		Return(ipBlockList, WithResponse(200, WithBody(ipBlockList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(IpBlockTables, "get ip-blocks").
		Return(nil)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllIpBlocksKeycloakFailure(test_framework *testing.T) {
	ipBlockList := []ipapisdk.IpBlock{}
	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksGet().
		Return(ipBlockList, nil, testutil.TestKeycloakError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestKeycloakError, err)
}

func TestGetAllIpBlocksPrinterFailure(test_framework *testing.T) {
	ipBlockList := ipmodels.GenerateIpBlockSdkList(2)

	var ipBlockTables []interface{}

	for _, ipBlock := range ipBlockList {
		ipBlockTables = append(ipBlockTables, tables.ToIpBlockTable(ipBlock))
	}

	PrepareIPMockClient(test_framework).
		IpBlocksGet().
		Return(ipBlockList, WithResponse(200, WithBody(ipBlockList)), nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(ipBlockTables, "get ip-blocks").
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
