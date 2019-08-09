package poweron

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

// ErrorCode to represent the status of the command execution.
// To be only used in testing.
// "OK"     => No errors.
// "ARGS"   => Arg amount error.
// "CLIENT" => Http Client failure error.
// "409"    => 409 response.
// "404"    => 404 response.
// "??"		=> Others
var ErrorCode = "500"

var P_OnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Powers on a specific server.",
	Long:  "Powers on a specific server.",
	Run: func(cmd *cobra.Command, args []string) {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-on':", len(args), "passed")
			ErrorCode = "ARGS"
			return
		}

		var resource = "servers/" + args[0] + "/actions/power-on"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering on server:", err)
			ErrorCode = "CLIENT"
			return
		}

		switch response.StatusCode {
		case 409:
			fmt.Println("Error: Conflict detected. Server is already powered-on.")
			ErrorCode = "409"
		case 404:
			fmt.Println("Error: Server with ID", args[0], "not found.")
			ErrorCode = "404"
		case 500:
			fmt.Println("Error: Internal server error. Please try again later.")
			ErrorCode = "500"
		case 200:
			fmt.Println("Powered on successfully.")
			ErrorCode = "OK"
		default:
			fmt.Println("Status:", response.Status)
			ErrorCode = "??"
		}
	},
}

func init() {
}
