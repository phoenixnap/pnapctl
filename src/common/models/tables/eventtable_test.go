package tables

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"phoenixnap.com/pnapctl/common/models"
	"phoenixnap.com/pnapctl/common/models/generators"
)

func TestToEventTable(test_framework *testing.T) {
	event := generators.GenerateEventSdk()
	table := ToEventTable(*event)

	assertEventsEqual(test_framework, *event, table)
}

func assertEventsEqual(test_framework *testing.T, event auditsdk.Event, table Event) {
	assert.Equal(test_framework, DerefString(event.Name), table.Name)
	assert.Equal(test_framework, event.Timestamp.String(), table.Timestamp)
	assert.Equal(test_framework, models.UserInfoToTableString(&event.UserInfo), table.UserInfo)
}
