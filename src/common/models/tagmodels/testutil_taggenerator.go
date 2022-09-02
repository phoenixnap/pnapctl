package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateTagSdk() *tagapisdk.Tag {
	values := []string{testutil.RandSeq(10)}

	return &tagapisdk.Tag{
		Id:          testutil.RandSeq(10),
		Name:        testutil.RandSeq(10),
		Values:      values,
		Description: testutil.RandSeqPointer(10),
		CreatedBy:   testutil.RandSeqPointer(5),
	}
}

func GenerateTagCreateCli() *TagCreate {
	return &TagCreate{
		Name:         testutil.RandSeq(10),
		Description:  testutil.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagUpdateCli() *TagUpdate {
	return &TagUpdate{
		Name:         testutil.RandSeq(10),
		Description:  testutil.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagListSdk(n int) []tagapisdk.Tag {
	var tagList []tagapisdk.Tag
	for i := 0; i < n; i++ {
		tagList = append(tagList, *GenerateTagSdk())
	}
	return tagList
}

func GenerateTagsDeleteResultSdk() *tagapisdk.DeleteResult {
	return &tagapisdk.DeleteResult{
		Result: testutil.RandSeq(10),
		TagId:  testutil.RandSeq(10),
	}
}

func GenerateResourceAssignmentSdk() *tagapisdk.ResourceAssignment {
	return &tagapisdk.ResourceAssignment{
		ResourceName: testutil.RandSeq(10),
		Value:        testutil.RandSeqPointer(10),
	}
}
