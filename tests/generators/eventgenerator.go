package generators

import (
	"time"

	auditapisdk "gitlab.com/phoenixnap/bare-metal-cloud/go-sdk.git/auditapi"
)

func GenerateEvent() auditapisdk.Event {
	return auditapisdk.Event{
		Name:      randSeqPointer(10),
		Timestamp: time.Now(),
		UserInfo: auditapisdk.UserInfo{
			AccountId: RandSeq(10),
			ClientId:  randSeqPointer(10),
			Username:  RandSeq(10),
		},
	}
}

func GenerateEvents(n int) []auditapisdk.Event {
	var eventList []auditapisdk.Event
	for i := 0; i < n; i++ {
		eventList = append(eventList, GenerateEvent())
	}
	return eventList
}
