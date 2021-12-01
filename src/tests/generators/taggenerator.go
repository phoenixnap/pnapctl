package generators

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

func GenerateTag() *tagapisdk.Tag {
	values := []string{RandSeq(10)}

	return &tagapisdk.Tag{
		Id:          RandSeq(10),
		Name:        RandSeq(10),
		Values:      &values,
		Description: RandSeqPointer(10),
	}
}

func GenerateTagCreate() *tagapisdk.TagCreate {
	return &tagapisdk.TagCreate{
		Name:         RandSeq(10),
		Description:  RandSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagUpdate() *tagapisdk.TagUpdate {
	return &tagapisdk.TagUpdate{
		Name:         RandSeq(10),
		Description:  RandSeqPointer(10),
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
		Result: RandSeq(10),
		TagId:  RandSeq(10),
	}
}
