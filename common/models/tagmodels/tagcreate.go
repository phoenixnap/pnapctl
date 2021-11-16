package tagmodels

import (
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
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
