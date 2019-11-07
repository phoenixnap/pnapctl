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

	// Marshalling errors: 04XX
	Marshalling          = "0400"
	MarshallingInPrinter = "0402"

	// File related errors:
	File             = "0500"
	FileReading      = "0503"
	FileDoesNotExist = "0504"

	// Miscellaneous errors: 99XX
	TablePrinterFailure = "9901"
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

// GenericFailedRequestError represents an error with performing a request.
// Requires the error that caused this issue and the command name being executed
func GenericFailedRequestError(err error, commandName string) error {
	if e, isCtlError := err.(Error); isCtlError {
		return e
	} else {
		return errors.New("Command '" + commandName + "' could not be performed. Please try again later.")
	}
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

// GenerateErrorIfNot200 returns an error if the response code is not 200.
// Ideally we want to use the response returned to us by the server but we return ageneric error if
// (i) there is no body
// (ii) The body can't be read (GO error)
// (iii) we can't deserialize the response
func GenerateErrorIfNot200(response *http.Response, commandName string) error {
	statusCode := response.StatusCode

	if statusCode == 200 {
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
