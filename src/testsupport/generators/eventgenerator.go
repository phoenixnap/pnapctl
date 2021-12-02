package generators

import (
	"time"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"phoenixnap.com/pnapctl/common/models/auditmodels"
)

func GenerateEvent() auditapisdk.Event {
	return auditapisdk.Event{
		Name:      RandSeqPointer(10),
		Timestamp: time.Now(),
		UserInfo: auditapisdk.UserInfo{
			AccountId: RandSeq(10),
			ClientId:  RandSeqPointer(10),
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

func GenerateQueryParams() auditmodels.EventsGetQueryParams {
	now := reprocessTime(time.Now())
	return auditmodels.EventsGetQueryParams{
		From:     &now,
		To:       &now,
		Limit:    10,
		Order:    "ASC",
		Username: RandSeq(10),
		Verb:     "PUT",
		Uri:      RandSeq(10),
	}
}

func reprocessTime(t time.Time) time.Time {
	result, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return result
}
