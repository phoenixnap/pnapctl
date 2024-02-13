package ip_blocks

import (
	"testing"

	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/stretchr/testify/assert"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
	. "phoenixnap.com/pnapctl/testsupport/mockhelp"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func TestGetIpBlocksSuccess(test_framework *testing.T) {
	ipBlock := generators.Generate[ipapi.IpBlock]()
	tableIpBlock := tables.ToShortIpBlockTable(ipBlock)

	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(&ipBlock, nil)

	ExpectToPrintSuccess(test_framework, tableIpBlock)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.NoError(test_framework, err)
}

func TestGetIpBlocksClientFailure(test_framework *testing.T) {
	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(nil, testutil.TestError)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Expected error
	expectedErr := ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}

func TestGetIpBlocksPrinterFailure(test_framework *testing.T) {
	ipBlock := generators.Generate[ipapi.IpBlock]()
	tableIpBlock := tables.ToShortIpBlockTable(ipBlock)

	PrepareIPMockClient(test_framework).
		IpBlocksGetById(RESOURCEID).
		Return(&ipBlock, nil)

	expectedErr := ExpectToPrintFailure(test_framework, tableIpBlock)

	err := GetIpBlockCmd.RunE(GetIpBlockCmd, []string{RESOURCEID})

	// Assertions
	assert.EqualError(test_framework, err, expectedErr.Error())
}
