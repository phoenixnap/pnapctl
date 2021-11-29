package auditmodels

import (
	"errors"
	"time"

	auditapisdk "github.com/phoenixnap/go-sdk-bmc/auditapi"
)

type EventsGetQueryParams struct {
	From     *time.Time
	To       *time.Time
	Limit    int
	Order    string
	Username string
	Verb     string
	Uri      string
}

func NewEventsGetQueryParams(from string, to string, limit int, order string, username string, verb string, uri string) (*EventsGetQueryParams, error) {
	var fromTime *time.Time
	var toTime *time.Time
	var validOrder string
	var validVerb string

	if from != "" {
		ft, err := time.Parse(time.RFC3339, from)
		if err != nil {
			return nil, errors.New("'From' (" + from + ") is not a valid date.")
		}
		fromTime = &ft
	}
	if to != "" {
		tt, err := time.Parse(time.RFC3339, to)
		if err != nil {
			return nil, errors.New("'To' (" + to + ") is not a valid date.")
		}
		toTime = &tt
	}
	switch order {
	case "":
		break
	case "ASC":
		fallthrough
	case "DESC":
		validOrder = order
		break
	default:
		return nil, errors.New("Invalid Order '" + order + "'. Valid values: 'ASC', 'DESC'")
	}

	switch verb {
	case "":
		break
	case "POST":
		fallthrough
	case "PUT":
		fallthrough
	case "PATCH":
		fallthrough
	case "DELETE":
		validVerb = verb
	default:
		return nil, errors.New("Invalid Verb '" + verb + "'. Valid values: 'POST', 'PUT', 'PATCH', 'DELETE'")
	}

	return &EventsGetQueryParams{
		From:     fromTime,
		To:       toTime,
		Limit:    limit,
		Order:    validOrder,
		Username: username,
		Verb:     validVerb,
		Uri:      uri,
	}, nil
}

func (queries EventsGetQueryParams) AttachToRequest(request auditapisdk.ApiEventsGetRequest) *auditapisdk.ApiEventsGetRequest {

	if queries.From != nil {
		request = request.From(*queries.From)
	}
	if queries.To != nil {
		request = request.To(*queries.To)
	}
	if queries.Limit != 0 {
		request = request.Limit(int32(queries.Limit))
	}
	if queries.Order != "" {
		request = request.Order(queries.Order)
	}
	if queries.Username != "" {
		request = request.Username(queries.Username)
	}
	if queries.Verb != "" {
		request = request.Verb(queries.Verb)
	}
	if queries.Uri != "" {
		request = request.Uri(queries.Uri)
	}

	return &request
}
