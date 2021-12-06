package auditmodels

import (
	"time"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
	"phoenixnap.com/pnapctl/tests/generators"
)

func GenerateEventSdk() auditapisdk.Event {
	return auditapisdk.Event{
		Name:      generators.RandSeqPointer(10),
		Timestamp: time.Now(),
		UserInfo: auditapisdk.UserInfo{
			AccountId: generators.RandSeq(10),
			ClientId:  generators.RandSeqPointer(10),
			Username:  generators.RandSeq(10),
		},
	}
}

func GenerateEventListSdk(n int) []auditapisdk.Event {
	var eventList []auditapisdk.Event
	for i := 0; i < n; i++ {
		eventList = append(eventList, GenerateEventSdk())
	}
	return eventList
}

func GenerateQueryParamsCli() EventsGetQueryParams {
	now := reprocessTime(time.Now())
	return EventsGetQueryParams{
		From:     &now,
		To:       &now,
		Limit:    10,
		Order:    "ASC",
		Username: generators.RandSeq(10),
		Verb:     "PUT",
		Uri:      generators.RandSeq(10),
	}
}

func reprocessTime(t time.Time) time.Time {
	result, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return result
}
