package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"phoenixnap.com/pnap-cli/pnapctl/client"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// HandleClientResponse determines whether
// (i) request was executed at all, if not return a generic error
// (ii) request was executed successfully, if so unmarshall response, print it and don't return an error
// (iii) request was executed unsuccessfully, if so attempt to use the return body to output an error
func HandleClientResponse(response *http.Response, err error, commandName string) error {
	if response == nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return ctlerrors.GenericNonRequestError(ctlerrors.ResponseBodyReadFailure, commandName)
		}

		responseBody := client.ResponseBody{}
		err = json.Unmarshal(body, &responseBody)

		if err != nil {
			return ctlerrors.GenericNonRequestError(ctlerrors.UnmarshallingResponseBody, commandName)
		}

		fmt.Println(responseBody.Result)
		return nil
	} else {
		return ctlerrors.HandleResponseError(response, commandName)
	}
}
