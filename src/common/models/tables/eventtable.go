package tables

import (
	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
	"phoenixnap.com/pnapctl/common/models"
)

type Event struct {
	Name      string `header:"Name"`
	Timestamp string `header:"Timestamp"`
	UserInfo  string `header:"User Info"`
}

func ToEventTable(event auditsdk.Event) Event {
	return Event{
		Name:      DerefString(event.Name),
		Timestamp: event.Timestamp.String(),
		UserInfo:  models.UserInfoToTableString(&event.UserInfo),
	}
}
