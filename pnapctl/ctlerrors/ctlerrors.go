package ctlerrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/* Error functions.
   To use for declaring/constructing errors. */

// Represents an invalid number of arguments passed to a command.
func InvalidNumberOfArgs(expected int, actual int, commandName string) error {
	var sb strings.Builder
	sb.WriteString("Only ")
	sb.WriteString(string(expected))
	sb.WriteString(" argument")

	if expected != 1 {
		sb.WriteString("s")
	}

	sb.WriteString(" can be passed for '")
	sb.WriteString(commandName)
	sb.WriteString("': ")
	sb.WriteString(string(actual))
	sb.WriteString(" passed.")
	return errors.New(sb.String())
}

// Represents a failure to read the response body.
func ResponseBodyReadError(err error) error {
	return errors.New("Error while reading body from response: " + err.Error())
}

// Represents an error from the Printer.
func PrinterError(err error) error {
	return errors.New("Error while printing output: " + err.Error())
}

// An error triggered when a body was expected, yet nothing was found.
func ExpectedBodyError(statusCode int) error {
	return errors.New("Expected a body to be in the response - no body was found. Status code: " + string(statusCode))
}

// Represents an error that occurred during Unmarshalling
func UnmarshallingError(resource string, err error) error {
	return errors.New("Couldn't unmarshal JSON response into '" + resource + "':" + err.Error())
}

// A generic error used for generic cases.
func GenericError(message string, err error) error {
	return errors.New(message + ": " + err.Error())
}

/* Generic error constants.
   A generic error function for each command. */

var GetServerGenericError = func(err error) error { return GenericError("Error while retrieving server", err) }
var GetServersGenericError = func(err error) error { return GenericError("Error while retrieving servers", err) }
var PowerOnServerGenericError = func(err error) error { return GenericError("Error while powering on server", err) }
var PowerOffServerGenericError = func(err error) error { return GenericError("Error while powering off server", err) }
var ShutdownServerGenericError = func(err error) error { return GenericError("Error while shutting down server", err) }
var RebootServerGenericError = func(err error) error { return GenericError("Error while rebooting server", err) }

/* Error handling.
   Structs and functions/methods for error handling. */

type result struct {
	Msg200 string
	Msg404 string
}

type BMCError struct {
	Message          string
	ValidationErrors []string
}

func Result() result {
	return result{
		Msg200: "",
		Msg404: "404 NOT FOUND",
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
		return ExpectedBodyError(statusCode)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return ResponseBodyReadError(err)
	}

	bmcErr := BMCError{}
	err = json.Unmarshal(body, &bmcErr)

	if err != nil {
		return UnmarshallingError("BMCError", err)
	}

	return errors.New(bmcErr.String())
}
