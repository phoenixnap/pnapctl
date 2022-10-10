package utils

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/printer"
)

func is2xxSuccessful(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return true
	} else {
		return false
	}
}

func CheckForErrors(httpResponse *http.Response, err error, commandName string) *error {
	var generatedError error = nil
	if httpResponse != nil && !is2xxSuccessful(httpResponse.StatusCode) {
		generatedError = ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		generatedError = ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	}

	return &generatedError
}

func SetupOutputFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

func SetupFullFlag(cmd *cobra.Command, full *bool, resource string) {
	cmd.Flags().BoolVar(full, "full", false, fmt.Sprintf("Shows all %s details", resource))
}

type Action string

const (
	CREATION Action = "creation"
	UPDATING Action = "updating"
)

func SetupFilenameFlag(cmd *cobra.Command, filename *string, action Action) {
	cmd.Flags().StringVarP(filename, "filename", "f", "", "File containing required information for "+string(action))
	cmd.MarkFlagRequired("filename")
}
