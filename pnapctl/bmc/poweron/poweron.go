package poweron

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

var P_OnCmd = &cobra.Command{
	Use:           "power-on",
	Short:         "Powers on a specific server.",
	Long:          "Powers on a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-on':", len(args), "passed")
			return errors.New("args")
		}

		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering on server:", err)
			return errors.New("client-fail")
		}

		switch response.StatusCode {
		case 409:
			fmt.Println("Error: Conflict detected. Server is already powered-on.")
			return errors.New("409")
		case 404:
			fmt.Println("Error: Server with ID", args[0], "not found.")
			return errors.New("404")
		case 500:
			fmt.Println("Error: Internal server error. Please try again later.")
			return errors.New("500")
		case 200:
			fmt.Println("Powered on successfully.")
			return nil
		default:
			fmt.Println("Status:", response.Status)
			return errors.New("p-off-generic")
		}
	},
}

func init() {
}
