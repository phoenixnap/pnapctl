package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type TagUpdate struct {
	Name         string  `json:"name" yaml:"name"`
	Description  *string `json:"description" yaml:"description"`
	IsBillingTag bool    `json:"isBillingTag" yaml:"isBillingTag"`
}

func (tagUpdate *TagUpdate) ToSdk() tagapisdk.TagUpdate {
	return tagapisdk.TagUpdate{
		Name:         tagUpdate.Name,
		Description:  tagUpdate.Description,
		IsBillingTag: tagUpdate.IsBillingTag,
	}
}

func TagUpdateFromSdk(tagUpdate *tagapisdk.TagUpdate) *TagUpdate {
	if tagUpdate == nil {
		return nil
	}

	return &TagUpdate{
		Name:         tagUpdate.Name,
		Description:  tagUpdate.Description,
		IsBillingTag: tagUpdate.IsBillingTag,
	}
}

func CreateTagUpdateFromFile(filename string, commandname string) (*tagapisdk.TagUpdate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var tagUpdate TagUpdate

	err = files.Unmarshal(data, &tagUpdate, commandname)

	if err != nil {
		return nil, err
	}

	sdkTagUpdate := tagUpdate.ToSdk()

	return &sdkTagUpdate, nil
}
