package tagmodels

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestTagCreateToSdk(test_framework *testing.T) {
	tagCreate := GenerateTagCreateCli()
	sdkTagCreate := tagCreate.ToSdk()

	assertEqualTagCreate(test_framework, *tagCreate, *sdkTagCreate)
}

// assertion functions
func assertEqualTagCreate(test_framework *testing.T, tagCreate TagCreate, sdkTagCreate tagapisdk.TagCreate) {
	assert.Equal(test_framework, tagCreate.Name, sdkTagCreate.Name)
	assert.Equal(test_framework, tagCreate.Description, sdkTagCreate.Description)
	assert.Equal(test_framework, tagCreate.IsBillingTag, sdkTagCreate.IsBillingTag)
}
