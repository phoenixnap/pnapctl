package tagmodels

import (
	tagapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/tagapi"
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
