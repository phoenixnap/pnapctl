package tables

import (
	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	auditmodels "phoenixnap.com/pnapctl/common/models/auditmodels"
)

type Event struct {
	Name      *string `header:"Name"`
	Timestamp string  `header:"Timestamp"`
	UserInfo  string  `header:"User Info"`
}

func ToEventTable(event auditsdk.Event) Event {
	return Event{
		Name:      event.Name,
		Timestamp: event.Timestamp.String(),
		UserInfo:  auditmodels.UserInfoToTableString(&event.UserInfo),
	}
}
