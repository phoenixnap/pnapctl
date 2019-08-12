package reboot

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"phoenixnap.com/pnap-cli/pnapctl/client"
)

var RebootCmd = &cobra.Command{
	Use:           "reboot",
	Short:         "Reboots a specific server.",
	Long:          "Reboots a specific server.",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			fmt.Println("only 1 argument can be passed for 'power-off':", len(args), "passed")
			return errors.New("args")
		}

		resource := "servers/" + args[0] + "/actions/reboot"
		response, err := client.MainClient.PerformPost(resource, bytes.NewBuffer([]byte{}))

		if err != nil {
			fmt.Println("Error while powering off server:", err)
			return errors.New("client-fail")
		}

		switch response.StatusCode {
		case 409:
			fmt.Println("Error: Conflict detected. Server can't be rebooted as it is powered off.")
			return errors.New("409")
		case 404:
			fmt.Println("Error: Server with ID", args[0], "not found.")
			return errors.New("404")
		case 500:
			fmt.Println("Error: Internal server error. Please try again later.")
			return errors.New("500")
		case 200:
			fmt.Println("Rebooted successfully.")
			return nil
		default:
			fmt.Println("Status:", response.Status)
			return errors.New("p-off-generic")
		}

		return nil
	},
}

func init() {}
