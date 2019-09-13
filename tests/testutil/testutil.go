package testutil

import (
	"errors"
	"strings"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// A generic test error.
var TestError = errors.New("TEST ERROR")

// An example of an error returned by the printer
const PrinterUnmarshalErrorMsg = "UnmarshallingInPrinter"

// An example of an error returned by the file processor.
const FileProcessorUnmarshalErrorMsg = "UnmarshallingInFileProcessor"

// Error returned by file processor when reading file.
var FileReadingError = errors.New("FileReading")

// A fake Error response from the server.
var GenericBMCError = ctlerrors.BMCError{
	Message:          "Something went wrong!",
	ValidationErrors: []string{},
}

// Asserting that two things are equal.
func AssertEqual(test_framework *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		test_framework.Error("\nExpected: '", expected, "'\nActual:   '", actual, "'")
	}
}

// Asserting that no error was found.
func AssertNoError(test_framework *testing.T, err error) {
	if err != nil {
		test_framework.Error("Expected no error, found: '", err, "'")
	}
}

// Asserting that an error has a specific errorcode.
func AssertErrorCode(test_framework *testing.T, err error, code string) {
	if !strings.Contains(err.Error(), code) {
		test_framework.Error("Errorcodes do not match. \nError: '", err, "'\nCode expected: '", code, "'")
	}
}
