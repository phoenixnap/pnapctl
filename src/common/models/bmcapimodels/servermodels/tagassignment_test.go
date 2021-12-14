package servermodels

import (
	"fmt"
	"testing"

	bmcapisdk "github.com/phoenixnap/go-sdk-bmc/bmcapi"
	"github.com/stretchr/testify/assert"
)

func TestTagAssignmentListFromSdk(test_framework *testing.T) {
	sdkModels := GenerateTagAssignmentListSdk(2)
	tagAssignmentList := TagAssignmentListFromSdk(&sdkModels)

	for i, tagAssignment := range *tagAssignmentList {
		assertEqualTagAssignment(test_framework, tagAssignment, sdkModels[i])
	}
}

func TestTagAssignmentListFromSdk_nilList(test_framework *testing.T) {
	assert.Nil(nil, TagAssignmentListFromSdk(nil))
}

func TestTagsToTableStrings_nilList(test_framework *testing.T) {
	result := TagsToTableStrings(nil)

	assert.Empty(test_framework, result)
}

func TestTagsToTableStrings(test_framework *testing.T) {
	sdkModels := GenerateTagAssignmentListSdk(1)
	result := TagsToTableStrings(&sdkModels)

	assert.Equal(test_framework, result[0], generateResultString(sdkModels[0]))
}

func assertEqualTagAssignment(test_framework *testing.T, cliTagAssignment TagAssignment, sdkTagAssignment bmcapisdk.TagAssignment) {
	assert.Equal(test_framework, cliTagAssignment.Id, sdkTagAssignment.Id)
	assert.Equal(test_framework, cliTagAssignment.Name, sdkTagAssignment.Name)
	assert.Equal(test_framework, cliTagAssignment.Value, sdkTagAssignment.Value)
	assert.Equal(test_framework, cliTagAssignment.IsBillingTag, sdkTagAssignment.IsBillingTag)

}

func generateResultString(tagAssignment bmcapisdk.TagAssignment) string {
	var tagValue = ": " + *tagAssignment.Value
	return fmt.Sprintf("(%s) %s%s", tagAssignment.Id, tagAssignment.Name, tagValue)
}
