package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestTagFromSdk(test_framework *testing.T) {
	tag := generators.GenerateTagSdk()
	table := TagFromSdk(*tag)

	assertTagsEqual(test_framework, *tag, table)
}

func assertTagsEqual(test_framework *testing.T, tag tagapisdk.Tag, table TagTable) {
	var resourceAssignments []string

	if tag.ResourceAssignments != nil {
		for _, x := range tag.ResourceAssignments {
			resourceAssignments = append(resourceAssignments, models.ResourceAssignmentToTableString(x))
		}
	}

	assert.Equal(test_framework, tag.Id, table.Id)
	assert.Equal(test_framework, tag.Name, table.Name)
	assert.Equal(test_framework, tag.Values, table.Values)
	assert.Equal(test_framework, DerefString(tag.Description), table.Description)
	assert.Equal(test_framework, tag.IsBillingTag, table.IsBillingTag)
	assert.Equal(test_framework, resourceAssignments, table.ResourceAssignments)
	assert.Equal(test_framework, DerefString(tag.CreatedBy), table.CreatedBy)
}
