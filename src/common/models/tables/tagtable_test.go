package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	tagapisdk "github.com/phoenixnap/go-sdk-bmc/tagapi/v3"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/utils/iterutils"
)

func TestTagFromSdk(test_framework *testing.T) {
	tag := generators.Generate[tagapisdk.Tag]()
	table := TagFromSdk(tag)

	assertTagsEqual(test_framework, tag, table)
}

func assertTagsEqual(test_framework *testing.T, tag tagapisdk.Tag, table TagTable) {
	resourceAssignments := iterutils.MapRef(tag.ResourceAssignments, models.ResourceAssignmentToTableString)

	assert.Equal(test_framework, tag.Id, table.Id)
	assert.Equal(test_framework, tag.Name, table.Name)
	assert.Equal(test_framework, tag.Values, table.Values)
	assert.Equal(test_framework, DerefString(tag.Description), table.Description)
	assert.Equal(test_framework, tag.IsBillingTag, table.IsBillingTag)
	assert.Equal(test_framework, resourceAssignments, table.ResourceAssignments)
	assert.Equal(test_framework, DerefString(tag.CreatedBy), table.CreatedBy)
}
