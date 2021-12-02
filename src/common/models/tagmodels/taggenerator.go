package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/testsupport/generators"
)

func GenerateTag() *tagapisdk.Tag {
	values := []string{generators.RandSeq(10)}

	return &tagapisdk.Tag{
		Id:          generators.RandSeq(10),
		Name:        generators.RandSeq(10),
		Values:      &values,
		Description: generators.RandSeqPointer(10),
	}
}

func GenerateCLITag() *Tag {
	values := []string{generators.RandSeq(10)}

	return &Tag{
		Id:          generators.RandSeq(10),
		Name:        generators.RandSeq(10),
		Values:      &values,
		Description: generators.RandSeqPointer(10),
	}
}

func GenerateTagCreate() *tagapisdk.TagCreate {
	return &tagapisdk.TagCreate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateCLITagCreate() *TagCreate {
	return &TagCreate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagUpdate() *tagapisdk.TagUpdate {
	return &tagapisdk.TagUpdate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateCLITagUpdate() *TagUpdate {
	return &TagUpdate{
		Name:         generators.RandSeq(10),
		Description:  generators.RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTags(n int) []tagapisdk.Tag {
	var tagList []tagapisdk.Tag
	for i := 0; i < n; i++ {
		tagList = append(tagList, *GenerateTag())
	}
	return tagList
}

func GenerateTagsDeleteResult() tagapisdk.DeleteResult {
	return tagapisdk.DeleteResult{
		Result: generators.RandSeq(10),
		TagId:  generators.RandSeq(10),
	}
}

func GenerateResourceAssignment() *tagapisdk.ResourceAssignment {
	return &tagapisdk.ResourceAssignment{
		ResourceName: generators.RandSeq(10),
		Value:        generators.RandSeqPointer(10),
	}
}

func GenerateCLIResourceAssignment() *ResourceAssignment {
	return &ResourceAssignment{
		ResourceName: generators.RandSeq(10),
		Value:        generators.RandSeqPointer(10),
	}
}
