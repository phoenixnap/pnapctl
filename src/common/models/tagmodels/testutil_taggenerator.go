package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/tests/generators"
)

func GenerateTagSdk() *tagapisdk.Tag {
	values := []string{generators.RandSeq(10)}

	return &tagapisdk.Tag{
		Id:          generators.RandSeq(10),
		Name:        generators.RandSeq(10),
		Values:      &values,
		Description: generators.RandSeqPointer(10),
	}
}

func GenerateTagCreateCli() *TagCreate {
	return &TagCreate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagUpdateCli() *TagUpdate {
	return &TagUpdate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
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

func GenerateTagsDeleteResultSdk() tagapisdk.DeleteResult {
	return tagapisdk.DeleteResult{
		Result: generators.RandSeq(10),
		TagId:  generators.RandSeq(10),
	}
}

func GenerateResourceAssignmentSdk() *tagapisdk.ResourceAssignment {
	return &tagapisdk.ResourceAssignment{
		ResourceName: generators.RandSeq(10),
		Value:        generators.RandSeqPointer(10),
	}
}
