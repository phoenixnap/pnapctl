package testutil

import (
	"errors"
	"testing"

	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

var TestError error = errors.New("TEST ERROR")
var GenericBMCError = ctlerrors.BMCError{
	Message:          "Something went wrong!",
	ValidationErrors: []string{},
}

func AssertEqual(test_framework *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		test_framework.Error("\nExpected: '", expected, "'\nActual:   '", actual, "'")
	}
}

func AssertNoError(test_framework *testing.T, err error) {
	if err != nil {
		test_framework.Error("Expected no error, found: '", err, "'")
	}
}
