package auditmodels

import (
	"time"

	auditsdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/auditapi"
)

type Event struct {
	Name      *string   `json:"name" yaml:"name"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	UserInfo  UserInfo  `json:"userInfo" yaml:"userInfo"`
}

func (event *Event) toSdk() *auditsdk.Event {
	return &auditsdk.Event{
		Name:      event.Name,
		Timestamp: event.Timestamp,
		UserInfo:  *event.UserInfo.toSdk(),
	}
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
