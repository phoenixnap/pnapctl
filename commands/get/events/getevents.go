package events

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/audit"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get events"

var ID string

var GetEventsCmd = &cobra.Command{
	Use:          "event",
	Short:        "Retrieve all events relating to your account.",
	Aliases:      []string{"quotas"},
	SilenceUsage: true,
	Long: `Retrieve all events relating to your account.
	
By default, the data is printed in table format.`,
	Example: `
# List all events in json format.
pnapctl get events -o json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getEvents()
	},
}

func getEvents() error {
	log.Debug("Getting events...")
	events, httpResponse, err := audit.Client.EventsGet()

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		return printer.PrintEventListResponse(events, commandName)
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

func init() {
	GetEventsCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
}
