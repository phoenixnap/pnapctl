package generators

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
)

func GenerateTag() *tagapisdk.Tag {
	values := []string{randSeq(10)}

	return &tagapisdk.Tag{
		Id:          randSeq(10),
		Name:        randSeq(10),
		Values:      &values,
		Description: randSeqPointer(10),
	}
}

func GenerateTagCreate() *tagapisdk.TagCreate {
	return &tagapisdk.TagCreate{
		Name:         randSeq(10),
		Description:  randSeqPointer(10),
		IsBillingTag: false,
	}
}

func GenerateTagUpdate() *tagapisdk.TagUpdate {
	return &tagapisdk.TagUpdate{
		Name:         randSeq(10),
		Description:  randSeqPointer(10),
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
		Result: randSeq(10),
		TagId:  randSeq(10),
	}
}
