package generators

import (
	ipapisdk "github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
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

func GenerateIpBlockTagSdk() ipapisdk.TagAssignmentRequest {
	return ipapisdk.TagAssignmentRequest{
		Name:  testutil.RandSeq(10),
		Value: testutil.RandSeqPointer(10),
	}
}

func GenerateIpBlockTagListSdk(n int) []ipapisdk.TagAssignmentRequest {
	var TagAssignmentRequestList []ipapisdk.TagAssignmentRequest

	for i := 0; i < n; i++ {
		TagAssignmentRequestList = append(TagAssignmentRequestList, GenerateIpBlockTagSdk())
	}

	return TagAssignmentRequestList
}

func GenerateDeleteIpBlockResultSdk() *ipapisdk.DeleteIpBlockResult {
	return &ipapisdk.DeleteIpBlockResult{
		Result:    testutil.RandSeq(10),
		IpBlockId: testutil.RandSeq(10),
	}
}

func GeneratePutTagIpBlockSdk() *ipapisdk.IpBlock {
	return &ipapisdk.IpBlock{
		Id:                   testutil.RandSeq(10),
		Location:             testutil.RandSeq(10),
		CidrBlockSize:        testutil.RandSeq(10),
		Cidr:                 testutil.RandSeq(10),
		Status:               testutil.RandSeq(10),
		AssignedResourceId:   testutil.RandSeqPointer(10),
		AssignedResourceType: testutil.RandSeqPointer(10),
	}
}

func GenerateIpBlockCreateSdk() ipapisdk.IpBlockCreate {
	return ipapisdk.IpBlockCreate{}
}

func GenerateIpBlockPatchSdk() ipapisdk.IpBlockPatch {
	return ipapisdk.IpBlockPatch{}
}
