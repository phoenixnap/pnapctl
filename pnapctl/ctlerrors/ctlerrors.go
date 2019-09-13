package ctlerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/* Client side errors. */

// FileNotExistError represents a file that does not exist
func FileNotExistError(filename string) error {
	return errors.New("The file '" + filename + "' does not exist.")
}

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
	File        = "0500"
	FileReading = "0503"

	// Miscellaneous errors: 99XX
	TablePrinterFailure = "9901"
)

/* Error functions.
   To use for declaring/constructing errors. */

// A generic error used for generic cases.
func GenericNonRequestError(errorCode string, command string) error {
	return errors.New("Command '" + command + "' has been performed, but something went wrong. Error code: " + errorCode)
}

// GenericFailedRequestError represents an error with performing a request.
func GenericFailedRequestError(command string) error {
	return errors.New("Command '" + command + "' could not be performed. Please try again later.")
}

/* Error handling.
   Structs and functions/methods for error handling. */

type BMCError struct {
	Message          string
	ValidationErrors []string
}

type result struct {
	Msg200      string
	Msg404      string
	CommandName string
}

func Result(commandName string) result {
	return result{
		Msg200:      "",
		Msg404:      "404 NOT FOUND",
		CommandName: commandName,
	}
}

func (r result) IfOk(message string) result {
	r.Msg200 = message
	return r
}

func (r result) IfNotFound(message string) result {
	r.Msg404 = message
	return r
}

func (b BMCError) String() string {
	if len(b.ValidationErrors) == 0 {
		return b.Message
	} else {
		return b.Message + "\n" + strings.Join(b.ValidationErrors, "\n")
	}
}

func (r result) UseResponse(response *http.Response) error {
	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println(r.Msg200)
		return nil
	} else if statusCode == 404 {
		return errors.New(r.Msg404)
	}

	if response.Body == nil {
		return GenericNonRequestError(r.CommandName, ExpectedBodyInErrorResponse)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return GenericNonRequestError(r.CommandName, ResponseBodyReadFailure)
	}

	bmcErr := BMCError{}
	err = json.Unmarshal(body, &bmcErr)

	if err != nil {
		return GenericNonRequestError(r.CommandName, UnmarshallingErrorBody)
	}

	return errors.New(bmcErr.String())
}
