package ctlerrors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

/*
A map of error codes.
Each errorcode has the structure XXYY,

	where XX refers to an error category,
	  and YY refers to a specific case.

If YY is 00, then the error is generic.
YY can also be categorized, for example currently:

	00 : General
	01 : In ctlerrors
	02 : In printer
	03 : In file processing
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

	// File processing errors: 05XX
	FileReading = "0503"

	// Flag parsing errors: 06XX
	InvalidFlagUsed = "0600"

	// Miscellaneous errors: 99XX
	TablePrinterFailure = "9901"
	// The error below typically happens either if there is a bug in the client or if the request body is incorrect
	ErrorSendingRequest = "9902"
)

/* Error functions.
   To use for declaring/constructing errors. */

// FileNotExistError represents a file that does not exist
func FileNotExistError(filename string) error {
	return CLIValidationError{
		Message: "The file '" + filename + "' does not exist.",
	}
}

// A generic error used for generic cases in commands.
func CreateCLIError(errorCode string, cause error) CLIError {
	return CLIError{
		Message: "Command '" + cmdname.CommandName + "' has been performed, but something went wrong. Error code: " + errorCode,
		Cause:   cause,
	}
}

// GenericFailedRequestError is used when an error occurs before the request has been executed.
// Requires the error that caused this issue, the command name being executed and a potential error code
func GenericFailedRequestError(err error, errorCode string) error {
	if e, isCtlError := err.(PnapctlError); isCtlError {
		return e
	}

	return CLIError{
		Message: "Command '" + cmdname.CommandName + "' could not be performed. Error code: " + errorCode,
		Cause:   err,
	}
}

// InvalidFlagValuePassedError is used when an invalid flag value is passed into the command.
// Accepts a list of accepted values to inform the user on what is allowed.
func InvalidFlagValuePassedError[T any](flagname string, flagvalue string, allowed []T) error {
	return CLIValidationError{
		Message: fmt.Sprintf("%s '%s' is invalid. Allowed values are %v", flagname, flagvalue, allowed),
	}
}

/* Error handling.
   Structs and functions/methods for error handling. */

// PnapctlError is an error that has been processed by ctlerrors.go and
// is ready to be shown to the user if so desired
type PnapctlError interface {
	Error() string
}

// Error that ocurrs on BMC API side (eg. validation, non 2xx response)
type BMCError struct {
	Message          string
	ValidationErrors []string
}

func (e BMCError) Error() string {
	if len(e.ValidationErrors) == 0 {
		return e.Message
	} else {
		return e.Message + "\n" + strings.Join(e.ValidationErrors, "\n")
	}
}

// Error that ocured at CLI level such as failing to unmarshal response, process response, timeout on authentication etc...
type CLIError struct {
	Message string
	Cause   error
}

func (e CLIError) Error() string {
	return e.Message
}

// Error to be used when some validation in CLI fails (eg. missing file, wrong format in file, wrong flag for command...)
type CLIValidationError struct {
	Message string
}

func (e CLIValidationError) Error() string {
	return e.Message
}

// HandleBMCError handles responses where the response is not 200.
// Ideally we want to use the response returned to us by the server but we return a generic error if
// (i) There is no response body (command executed but no body returned)
// (ii) The response body can't be read (probably GO error)
// (iii) we can't deserialize the response (probably a server error)
func HandleBMCError(response *http.Response) error {
	if response != nil && response.StatusCode == 200 {
		// Technically we should never enter here. If we do, something went wrong previously.
		return nil
	}

	if response.Body == nil {
		return CreateCLIError(ExpectedBodyInErrorResponse, nil)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return CreateCLIError(ResponseBodyReadFailure, err)
	}

	bmcErr := BMCError{}
	err = json.Unmarshal(body, &bmcErr)

	if err != nil || len(bmcErr.Error()) == 0 {
		return CreateCLIError(UnmarshallingErrorBody, err)
	}

	return bmcErr
}
