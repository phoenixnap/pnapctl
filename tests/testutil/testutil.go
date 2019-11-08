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
func AssertEqual(testFramework *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		testFramework.Error("\nExpected: '", expected, "'\nActual:   '", actual, "'")
	}
}

// AssertStringContains checks whether testString contains a particular string subset
func AssertStringContains(testFramework *testing.T, testString string, contains string) {
	if !strings.Contains(testString, contains) {
		testFramework.Error("\n", testString, " does not contain ", contains)
	}
}

// Asserting that no error was found.
func AssertNoError(testFramework *testing.T, err error) {
	if err != nil {
		testFramework.Error("Expected no error, found: '", err, "'")
	}
}

// Asserting that an error has a specific errorcode.
func AssertErrorCode(testFramework *testing.T, err error, code string) {
	if !strings.Contains(err.Error(), code) {
		testFramework.Error("Errorcodes do not match. \nError: '", err, "'\nCode expected: '", code, "'")
	}
}
