package generators

import (
	"time"

	"phoenixnap.com/pnapctl/common/models/queryparams/audit"
)

var GenerateQueryParamsSdk = Generator(func(event *audit.EventsGetQueryParams) {
	event.Order = "ASC"
	event.Verb = "POST"
	t := reprocessTime(*event.From)
	event.From = &t
	event.To = &t
})

func reprocessTime(t time.Time) time.Time {
	result, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return result
}
