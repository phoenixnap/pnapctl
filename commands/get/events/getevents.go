package events

import (
	"github.com/google/martian/log"
	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/common/client/audit"
	"phoenixnap.com/pnap-cli/common/ctlerrors"
	"phoenixnap.com/pnap-cli/common/models/auditmodels"
	"phoenixnap.com/pnap-cli/common/printer"
)

const commandName string = "get events"

var ID string

var GetEventsCmd = &cobra.Command{
	Use:          "events",
	Short:        "Retrieve all events relating to your account.",
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
	params, err := auditmodels.NewEventsGetQueryParams(From, To, Limit, Order, Username, Verb)
	if err != nil {
		return err
	}

	events, httpResponse, err := audit.Client.EventsGet(*params)

	if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else if httpResponse.StatusCode == 200 {
		return printer.PrintEventListResponse(events, commandName)
	} else {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	}
}

var From string
var To string
var Limit int
var Order string
var Username string
var Verb string

func init() {
	GetEventsCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetEventsCmd.PersistentFlags().StringVar(&From, "from", "", "A 'from' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetEventsCmd.PersistentFlags().StringVar(&To, "to", "", "A 'to' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetEventsCmd.PersistentFlags().IntVar(&Limit, "limit", 0, "Limit the number of records returned.")
	GetEventsCmd.PersistentFlags().StringVar(&Order, "order", "", "Ordering of the event's time. Must be 'ASC' or 'DESC'")
	GetEventsCmd.PersistentFlags().StringVar(&Username, "username", "", "The username that did the actions.")
	GetEventsCmd.PersistentFlags().StringVar(&Verb, "verb", "", "The HTTP verb corresponding to the action. Must be 'POST', 'PUT', 'PATCH', 'DELETE'")
}
