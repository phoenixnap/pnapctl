package events

import (
	"github.com/spf13/cobra"
	"phoenixnap.com/pnapctl/common/client/audit"
	"phoenixnap.com/pnapctl/common/ctlerrors"
	"phoenixnap.com/pnapctl/common/models/auditmodels"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
)

const commandName string = "get events"

var ID string

var GetEventsCmd = &cobra.Command{
	Use:          "event",
	Short:        "Retrieve all events relating to your account.",
	Aliases:      []string{"events"},
	SilenceUsage: true,
	Long: `Retrieve all events relating to your account.
	
By default, the data is printed in table format.`,
	Example: `
# List all events.
pnapctl get events [--from <FROM>] [--to <TO>] [--limit <LIMIT>] [--order <ORDER>] [--username <USERNAME>] [--verb <VERB>] [--uri <URI>] [--output <OUTPUT_TYPE>]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return getEvents()
	},
}

func getEvents() error {
	params, err := auditmodels.NewEventsGetQueryParams(From, To, Limit, Order, Username, Verb, Uri)
	if err != nil {
		return err
	}

	events, httpResponse, err := audit.Client.EventsGet(*params)

	if httpResponse != nil && !utils.Is2xxSuccessful(httpResponse.StatusCode) {
		return ctlerrors.HandleBMCError(httpResponse, commandName)
	} else if err != nil {
		return ctlerrors.GenericFailedRequestError(err, commandName, ctlerrors.ErrorSendingRequest)
	} else {
		return printer.PrintEventListResponse(events, commandName)
	}
}

var From string
var To string
var Limit int
var Order string
var Username string
var Verb string
var Uri string

func init() {
	GetEventsCmd.PersistentFlags().StringVarP(&printer.OutputFormat, "output", "o", "table", "Define the output format. Possible values: table, json, yaml")
	GetEventsCmd.PersistentFlags().StringVar(&From, "from", "", "A 'from' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetEventsCmd.PersistentFlags().StringVar(&To, "to", "", "A 'to' filter. Needs to be in the following format: '2021-04-27T16:24:57.123Z'")
	GetEventsCmd.PersistentFlags().IntVar(&Limit, "limit", 0, "Limit the number of records returned.")
	GetEventsCmd.PersistentFlags().StringVar(&Order, "order", "", "Ordering of the event's time. Must be 'ASC' or 'DESC'")
	GetEventsCmd.PersistentFlags().StringVar(&Username, "username", "", "The username that did the actions.")
	GetEventsCmd.PersistentFlags().StringVar(&Verb, "verb", "", "The HTTP verb corresponding to the action. Must be 'POST', 'PUT', 'PATCH', 'DELETE'")
	GetEventsCmd.PersistentFlags().StringVar(&Uri, "uri", "", "The request URI.")
}
