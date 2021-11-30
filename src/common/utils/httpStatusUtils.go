package utils

import (
	"net/http"

	"phoenixnap.com/pnapctl/common/ctlerrors"
)

func Is2xxSuccessful(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return true
	} else {
		return false
	}
}

func CheckForErrors(httpResponse *http.Response, err error, commandName string) *error {
	var generatedError error = nil
	if httpResponse != nil && !Is2xxSuccessful(httpResponse.StatusCode) {
		generatedError = ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		generatedError = ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	}

	return &generatedError
}
