package auditmodels

import (
	"time"

	auditsdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
)

type Event struct {
	Name      *string   `json:"name" yaml:"name"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	UserInfo  UserInfo  `json:"userInfo" yaml:"userInfo"`
}

func EventFromSdk(event *auditsdk.Event) *Event {
	if event == nil {
		return nil
	}

	return &Event{
		Name:      event.Name,
		Timestamp: event.Timestamp,
		UserInfo:  *UserInfoFromSdk(&event.UserInfo),
	}
}
