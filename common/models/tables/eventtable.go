package tables

import (
	"time"

	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	auditmodels "phoenixnap.com/pnap-cli/common/models/auditmodels"
)

type Event struct {
	Name      *string   `header:"Name"`
	Timestamp time.Time `header:"Timestamp"`
	UserInfo  string    `header:"User Info"`
}

func ToEventTable(event auditsdk.Event) Event {
	return Event{
		Name:      event.Name,
		Timestamp: event.Timestamp,
		UserInfo:  auditmodels.UserInfoToTableString(&event.UserInfo),
	}
}
