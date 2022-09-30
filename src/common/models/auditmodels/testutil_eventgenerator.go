package auditmodels

import (
	"time"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

func GenerateEventSdk() *auditapisdk.Event {
	return &auditapisdk.Event{
		Name:      testutil.RandSeqPointer(10),
		Timestamp: time.Now(),
		UserInfo:  GenerateUserInfoSdk(),
	}
}

func GenerateEventListSdk(n int) []auditapisdk.Event {
	var eventList []auditapisdk.Event
	for i := 0; i < n; i++ {
		eventList = append(eventList, *GenerateEventSdk())
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
		Username: testutil.RandSeq(10),
		Verb:     "PUT",
		Uri:      testutil.RandSeq(10),
	}
}

func GenerateUserInfoSdk() auditapisdk.UserInfo {
	return auditapisdk.UserInfo{
		AccountId: testutil.RandSeq(10),
		ClientId:  testutil.RandSeqPointer(10),
		Username:  testutil.RandSeq(10),
	}
}

func reprocessTime(t time.Time) time.Time {
	result, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return result
}
