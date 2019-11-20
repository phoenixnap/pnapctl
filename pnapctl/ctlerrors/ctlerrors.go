package ctlerrors

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

/*	A map of error codes.
	Each errorcode has the structure XXYY,
		where XX refers to an error category,
		  and YY refers to a specific case.

	If YY is 00, then the error is generic.
	YY can also be categorized, for example currently:
		00 : General
		01 : In ctlerrors
		02 : In printer
*/
const (
	// Reading Body Failure errors: 01XX
	ResponseBodyReadFailure = "0100"

	// Expected Body errors: 02XX
	ExpectedBodyInResponse      = "0200"
	ExpectedBodyInErrorResponse = "0201"

	// Unmarshalling errors: 03XX
	Unmarshalling                = "0300"
	UnmarshallingErrorBody       = "0301"
	UnmarshallingInPrinter       = "0302"
	UnmarshallingInFileProcessor = "0303"
	UnmarshallingResponseBody    = "0304"

	// Marshalling errors: 04XX
	Marshalling          = "0400"
	MarshallingInPrinter = "0402"

	// File related errors:
	File             = "0500"
	FileReading      = "0503"
	FileDoesNotExist = "0504"

	// Miscellaneous errors: 99XX
	TablePrinterFailure = "9901"
	// The error below typically happens either if there is a bug in the client or if the request body is incorrect
	ErrorSendingRequest = "9902"
)

/* Error functions.
   To use for declaring/constructing errors. */

// FileNotExistError represents a file that does not exist
func FileNotExistError(filename string) error {
	return errors.New("The file '" + filename + "' does not exist.")
}

// A generic error used for generic cases.
func GenericNonRequestError(errorCode string, command string) error {
	return errors.New("Command '" + command + "' has been performed, but something went wrong. Error code: " + errorCode)
}

// GenericFailedRequestError is used when an error occurs before the request has been executed.
// Requires the error that caused this issue, the command name being executed and a potential error code
func GenericFailedRequestError(err error, commandName string, errorCode string) error {
	if e, isCtlError := err.(Error); isCtlError {
		return e
	}

	return errors.New("Command '" + commandName + "' could not be performed. Error code: " + errorCode)
}

/* Error handling.
   Structs and functions/methods for error handling. */

// Error is an error that has been processed by ctlerrors.go and
// is ready to be shown to the user if so desired
type Error struct {
	Msg   string
	Cause error
}

func (e Error) Error() string {
	return e.Msg
}

type BMCError struct {
	Message          string
	ValidationErrors []string
}

func (b BMCError) String() string {
	if len(b.ValidationErrors) == 0 {
		return b.Message
	} else {
		return b.Message + "\n" + strings.Join(b.ValidationErrors, "\n")
	}
}

// HandleResponseError handles responses where the response is not 200.
// Ideally we want to use the response returned to us by the server but we return a generic error if
// (i) There is no response body (command executed but no body returned)
// (ii) The response body can't be read (probably GO error)
// (iii) we can't deserialize the response (probably a server error)
func HandleResponseError(response *http.Response, commandName string) error {
	if response != nil && response.StatusCode == 200 {
		// Technically we should never enter here. If we do, something went wrong previously.
		return nil
	}

	if response.Body == nil {
		return GenericNonRequestError(ExpectedBodyInErrorResponse, commandName)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return GenericNonRequestError(ResponseBodyReadFailure, commandName)
	}

	bmcErr := BMCError{}
	err = json.Unmarshal(body, &bmcErr)

	if err != nil || len(bmcErr.String()) == 0 {
		return GenericNonRequestError(UnmarshallingErrorBody, commandName)
	}

	return errors.New(bmcErr.String())
}
