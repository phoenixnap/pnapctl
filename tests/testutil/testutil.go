package testutil

import (
	"errors"
	"strings"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// A generic test error.
var TestError = errors.New("TEST ERROR")
var TestKeycloakError = ctlerrors.Error{Msg: "Failed to resolve provided credentials", Cause: TestError}

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
