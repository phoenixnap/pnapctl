package utils

import (
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/printer"
)

func SetupOutputFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}

func SetupFullFlag(cmd *cobra.Command, full *bool, resource string) {
	cmd.Flags().BoolVar(full, "full", false, fmt.Sprintf("Shows all %s details", resource))
}

type Action string

const (
	CREATION    Action = "creation"
	UPDATING    Action = "updating"
	RESERVATION Action = "reservation"
	TAGGING     Action = "tagging"
	CONVERSION  Action = "conversion"
	DELETION    Action = "deletion"
	DEPROVISION Action = "deprovisioning"
	SUBMISSION  Action = "submission"
	PROVISION   Action = "provisioning"
)

func SetupFilenameFlag(cmd *cobra.Command, filename *string, action Action) {
	cmd.Flags().StringVarP(filename, "filename", "f", "", "File containing required information for "+string(action))
	cmd.MarkFlagRequired("filename")
}
