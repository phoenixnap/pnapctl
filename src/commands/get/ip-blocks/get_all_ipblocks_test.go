package ip_blocks

import (
	"errors"
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetAllIpBlocksSuccess(test_framework *testing.T) {
	ipBlockList := testutil.GenN(2, generators.Generate[ipapi.IpBlock])
	IpBlockTables := iterutils.MapInterface(ipBlockList, tables.ToShortIpBlockTable)

	// Mocking
	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(ipBlockList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(IpBlockTables).
		Return(nil)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetAllIpBlocksClientFailure(test_framework *testing.T) {
	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(nil, testutil.TestError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Equal(test_framework, testutil.TestError, err)
}

func TestGetAllIpBlocksPrinterFailure(test_framework *testing.T) {
	ipBlockList := testutil.GenN(2, generators.Generate[ipapi.IpBlock])
	ipBlockTables := iterutils.MapInterface(ipBlockList, tables.ToShortIpBlockTable)

	PrepareIPMockClient(test_framework).
		IpBlocksGet(tags).
		Return(ipBlockList, nil)

	PrepareMockPrinter(test_framework).
		PrintOutput(ipBlockTables).
		Return(errors.New(ctlerrors.UnmarshallingInPrinter))

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{})

	// Assertions
	assert.Contains(test_framework, err.Error(), ctlerrors.UnmarshallingInPrinter)
}
