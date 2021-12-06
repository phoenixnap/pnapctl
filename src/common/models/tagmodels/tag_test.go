package tagmodels

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
)

// tests
func TestTagFromSdk(test_framework *testing.T) {
	sdkTag := GenerateTagSdk()
	tag := TagFromSdk(sdkTag)

	assertEqualTag(test_framework, *tag, *sdkTag)
}

func TestTagWithResourceAssignmentsFromSdk(test_framework *testing.T) {
	sdkTag := GenerateTagSdk()
	sdkTag.ResourceAssignments = &[]tagapisdk.ResourceAssignment{*GenerateResourceAssignmentSdk()}
	tag := TagFromSdk(sdkTag)

	assertEqualTag(test_framework, *tag, *sdkTag)
}

func TestNilTagFromSdk(test_framework *testing.T) {
	tag := TagFromSdk(nil)

	assert.Nil(test_framework, tag)
}

// assertion functions
func assertEqualTag(test_framework *testing.T, tag Tag, sdkTag tagapisdk.Tag) {
	assert.Equal(test_framework, tag.Id, sdkTag.Id)
	assert.Equal(test_framework, tag.Name, sdkTag.Name)
	assert.Equal(test_framework, tag.Values, sdkTag.Values)
	assert.Equal(test_framework, tag.Description, sdkTag.Description)
	assert.Equal(test_framework, tag.IsBillingTag, sdkTag.IsBillingTag)

	if tag.ResourceAssignments == nil {
		assert.Nil(test_framework, sdkTag.ResourceAssignments, "CLI Tag's Resource Assignments are nil, but not SDK Tag's Resource Assignments.")
	} else if sdkTag.ResourceAssignments == nil {
		assert.Nil(test_framework, tag.ResourceAssignments, "SDK Tag's Resource Assignments are nil, but not CLI Tag's Resource Assignments.")
	} else {
		assert.Equal(test_framework, len(*tag.ResourceAssignments), len(*sdkTag.ResourceAssignments))
		for i, _ := range *tag.ResourceAssignments {
			assertEqualResourceAssignment(test_framework, (*tag.ResourceAssignments)[i], (*sdkTag.ResourceAssignments)[i])
		}
	}
}
