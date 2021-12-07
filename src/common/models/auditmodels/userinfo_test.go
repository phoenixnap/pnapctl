package auditmodels

import (
	"fmt"
	"testing"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestUserInfoFromSdk(test_framework *testing.T) {
	sdkUserInfo := GenerateUserInfoSdk()
	userInfo := UserInfoFromSdk(&sdkUserInfo)

	assertEqualUserInfo(test_framework, *userInfo, sdkUserInfo)
}

func TestNilUserInfoFromSdk(test_framework *testing.T) {
	userInfo := UserInfoFromSdk(nil)

	assert.Nil(test_framework, userInfo)
}

func TestUserInfoToTableString(test_framework *testing.T) {
	sdkUserInfo := GenerateUserInfoSdk()

	assert.Equal(test_framework, fmt.Sprintf("Account:(%s)\nClientId:(%s)\nUsername:%s", sdkUserInfo.AccountId, *sdkUserInfo.ClientId, sdkUserInfo.Username), UserInfoToTableString(&sdkUserInfo))
}

func TestNillUserInfoToTableString_noRequests(test_framework *testing.T) {
	assert.Equal(test_framework, "", UserInfoToTableString(nil))
}

// assertion functions
func assertEqualUserInfo(test_framework *testing.T, userInfo UserInfo, sdkUserInfo auditapisdk.UserInfo) {
	assert.Equal(test_framework, userInfo.AccountId, sdkUserInfo.AccountId)
	assert.Equal(test_framework, userInfo.ClientId, sdkUserInfo.ClientId)
	assert.Equal(test_framework, userInfo.Username, sdkUserInfo.Username)
}
