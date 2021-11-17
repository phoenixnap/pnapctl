package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	files "phoenixnap.com/pnap-cli/common/fileprocessor"
)

type TagCreate struct {
	Name         string  `json:"name" yaml:"name"`
	Description  *string `json:"description" yaml:"description"`
	IsBillingTag bool    `json:"isBillingTag" yaml:"isBillingTag"`
}

func (tagCreate *TagCreate) ToSdk() tagapisdk.TagCreate {
	return tagapisdk.TagCreate{
		Name:         tagCreate.Name,
		Description:  tagCreate.Description,
		IsBillingTag: tagCreate.IsBillingTag,
	}
}

func CreateTagCreateFromFile(filename string, commandname string) (*tagapisdk.TagCreate, error) {
	files.ExpandPath(&filename)

	data, err := files.ReadFile(filename, commandname)

	if err != nil {
		return nil, err
	}

	var tagCreate TagCreate

	err = files.Unmarshal(data, &tagCreate, commandname)

	if err != nil {
		return nil, err
	}

	sdkTagCreate := tagCreate.ToSdk()

	return &sdkTagCreate, nil
}
