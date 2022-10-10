package auditmodels

import (
	"fmt"

	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
)

type UserInfo struct {
	AccountId string  `json:"accountId" yaml:"accountId"`
	ClientId  *string `json:"clientId" yaml:"clientId"`
	Username  string  `json:"username" yaml:"username"`
}

func UserInfoFromSdk(userInfo *auditsdk.UserInfo) *UserInfo {
	if userInfo == nil {
		return nil
	}

	return &UserInfo{
		AccountId: userInfo.AccountId,
		ClientId:  userInfo.ClientId,
		Username:  userInfo.Username,
	}
}

func UserInfoToTableString(userInfo *auditsdk.UserInfo) string {
	if userInfo == nil {
		return ""
	}

	return fmt.Sprintf("Account:(%s)\nClientId:(%s)\nUsername:%s", userInfo.AccountId, *userInfo.ClientId, userInfo.Username)
}
