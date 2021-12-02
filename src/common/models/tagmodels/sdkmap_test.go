package tagmodels

import (
	"testing"

	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"github.com/stretchr/testify/assert"
)

//tests
func TestTagFromSdk(test_framework *testing.T) {
	sdkTag := GenerateTag()
	tag := TagFromSdk(sdkTag)

	assertEqualTag(test_framework, *tag, *sdkTag)
}

func TestTagWithResourceAssignmentsFromSdk(test_framework *testing.T) {
	sdkTag := GenerateTag()
	sdkTag.ResourceAssignments = &[]tagapisdk.ResourceAssignment{*GenerateResourceAssignment()}
	tag := TagFromSdk(sdkTag)

	assertEqualTag(test_framework, *tag, *sdkTag)
}

func TestNilTagFromSdk(test_framework *testing.T) {
	tag := TagFromSdk(nil)

	assert.Nil(test_framework, tag)
}

func TestTagCreateToSdk(test_framework *testing.T) {
	tagCreate := GenerateCLITagCreate()
	sdkTagCreate := tagCreate.ToSdk()

	assertEqualTagCreate(test_framework, *tagCreate, *sdkTagCreate)
}

func TestTagUpdateToSdk(test_framework *testing.T) {
	tagUpdate := GenerateCLITagUpdate()
	sdkTagUpdate := tagUpdate.ToSdk()

	assertEqualTagUpdate(test_framework, *tagUpdate, *sdkTagUpdate)
}

func TestResourceAssignmentFromSdk(test_framework *testing.T) {
	sdkResourceAssignment := GenerateResourceAssignment()
	resourceAssignment := ResourceAssignmentFromSdk(sdkResourceAssignment)

	assertEqualResourceAssignment(test_framework, *resourceAssignment, *sdkResourceAssignment)
}

func TestNilResourceAssignmentFromSdk(test_framework *testing.T) {
	resourceAssignment := ResourceAssignmentFromSdk(nil)

	assert.Nil(test_framework, resourceAssignment)
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

func assertEqualTagCreate(test_framework *testing.T, tagCreate TagCreate, sdkTagCreate tagapisdk.TagCreate) {
	assert.Equal(test_framework, tagCreate.Name, sdkTagCreate.Name)
	assert.Equal(test_framework, tagCreate.Description, sdkTagCreate.Description)
	assert.Equal(test_framework, tagCreate.IsBillingTag, sdkTagCreate.IsBillingTag)
}

func assertEqualTagUpdate(test_framework *testing.T, tagUpdate TagUpdate, sdkTagUpdate tagapisdk.TagUpdate) {
	assert.Equal(test_framework, tagUpdate.Name, sdkTagUpdate.Name)
	assert.Equal(test_framework, tagUpdate.Description, sdkTagUpdate.Description)
	assert.Equal(test_framework, tagUpdate.IsBillingTag, sdkTagUpdate.IsBillingTag)
}

func assertEqualResourceAssignment(test_framework *testing.T, r1 ResourceAssignment, r2 tagapisdk.ResourceAssignment) {
	assert.Equal(test_framework, r1.ResourceName, r2.ResourceName)
	assert.Equal(test_framework, r1.Value, r2.Value)
}
