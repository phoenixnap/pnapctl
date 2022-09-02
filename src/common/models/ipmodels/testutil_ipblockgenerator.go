package ipmodels

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateIpBlockSdk() ipapisdk.IpBlock {
	return ipapisdk.IpBlock{
		Id:                   testutil.RandSeq(10),
		Location:             testutil.RandSeq(10),
		CidrBlockSize:        testutil.RandSeq(10),
		Cidr:                 testutil.RandSeq(10),
		Status:               testutil.RandSeq(10),
		AssignedResourceId:   testutil.RandSeqPointer(10),
		AssignedResourceType: testutil.RandSeqPointer(10),
	}
}

func GenerateIpBlockSdkList(n int) []ipapisdk.IpBlock {
	var ipBlockList []ipapisdk.IpBlock
	for i := 0; i < n; i++ {
		ipBlockList = append(ipBlockList, GenerateIpBlockSdk())
	}
	return ipBlockList
}

func GenerateIpBlockCreateCLI() IpBlockCreate {
	return IpBlockCreate{
		CidrBlockSize: testutil.RandSeq(10),
		Location:      testutil.RandSeq(10),
	}
}

func GenerateIpBlockPatchCLI() IpBlockPatch {
	return IpBlockPatch{
		Description: testutil.RandSeqPointer(10),
	}
}

func GenerateDeleteIpBlockResultSdk() *ipapisdk.DeleteIpBlockResult {
	return &ipapisdk.DeleteIpBlockResult{
		Result:    testutil.RandSeq(10),
		IpBlockId: testutil.RandSeq(10),
	}
}
