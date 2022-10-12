package client

import (
	"net/http"
	"time"

	"phoenixnap.com/pnapctl/common/ctlerrors"
)

func IsZeroValue[T comparable](item T) bool {
	var t T
	return t == item
}

func ParseDate(item string) *time.Time {
	ft, err := time.Parse(time.RFC3339, item)
	if err != nil {
		return nil
	}
	return &ft
}

func HandleResponse[T any](response T, httpResponse *http.Response, err error) (T, error) {
	return response, HandleResponseWithoutBody(httpResponse, err)
}

func HandleResponseWithoutBody(httpResponse *http.Response, err error) error {
	if httpResponse != nil && httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300 {
		return ctlerrors.HandleBMCError(httpResponse)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, ctlerrors.ErrorSendingRequest)
	}
	return nil
}
