package poweroff

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

var P_OffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Powers off a specific server.",
	Long:  "Powers off a specific server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// If more than one argument is passed, report error and panic.
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			return errors.New("args")
		}

		var resource = "servers/" + args[0] + "/actions/power-off"
		var response, err = client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			// Generic error with PerformPost
			fmt.Println("Error while powering off server:", err)
			return errors.New("client-fail")
		}

		switch response.StatusCode {
		case 409:
			fmt.Println("Error: Conflict detected. Server is already powered-off.")
			return errors.New("409")
		case 404:
			fmt.Println("Error: Server with ID", args[0], "not found.")
			return errors.New("404")
		case 500:
			fmt.Println("Error: Internal server error. Please try again later.")
			return errors.New("500")
		case 200:
			fmt.Println("Powered off successfully.")
			return nil
		default:
			fmt.Println("Status:", response.Status)
			return errors.New("p-off-generic")
		}
	},
}

func init() {
}
