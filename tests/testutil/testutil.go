package testutil

import (
	"errors"

	"phoenixnap.com/pnap-cli/common/ctlerrors"
)

// A generic test error.
var TestError = errors.New("TEST ERROR")
var TestKeycloakError = ctlerrors.CLIError{Message: "Failed to resolve provided credentials", Cause: TestError}

// A fake Error response from the server.
var GenericBMCError = ctlerrors.BMCError{
	Message:          "Something went wrong!",
	ValidationErrors: []string{},
}
