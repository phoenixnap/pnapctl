package tagmodels

import (
	tagapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/tagapi"
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
