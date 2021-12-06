package tagmodels

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestTagUpdateToSdk(test_framework *testing.T) {
	tagUpdate := GenerateTagUpdateCli()
	sdkTagUpdate := tagUpdate.ToSdk()

	assertEqualTagUpdate(test_framework, *tagUpdate, *sdkTagUpdate)
}

// assertion functions
func assertEqualTagUpdate(test_framework *testing.T, tagUpdate TagUpdate, sdkTagUpdate tagapisdk.TagUpdate) {
	assert.Equal(test_framework, tagUpdate.Name, sdkTagUpdate.Name)
	assert.Equal(test_framework, tagUpdate.Description, sdkTagUpdate.Description)
	assert.Equal(test_framework, tagUpdate.IsBillingTag, sdkTagUpdate.IsBillingTag)
}
