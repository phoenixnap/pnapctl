package locations

import (
	"github.com/spf13/cobra"
	"github.com/rs/zerolog/log"

	"phoenixnap.com/pnapctl/common/client/locations"
	"phoenixnap.com/pnapctl/common/printer"
	"phoenixnap.com/pnapctl/common/utils"
	"phoenixnap.com/pnapctl/common/utils/cmdname"
)

var (
	Location        string
	ProductCategory string
)

func init() {
	utils.SetupOutputFlag(GetLocationsCmd)

	GetLocationsCmd.PersistentFlags().StringVar(&Location, "location", "", "Location to filter by.")
	GetLocationsCmd.PersistentFlags().StringVar(&ProductCategory, "product-category", "", "Product category to filter locations by.")
}

var GetLocationsCmd = &cobra.Command{
	Use:          "location",
	Short:        "Retrieves all locations.",
	Aliases:      []string{"locations"},
	SilenceUsage: true,
	Long: `Retrieve all locations.
	
Prints all information about locations.
By default, the data is printed in table format.`,
	Example: ``,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmdname.SetCommandName(cmd)
		return getLocations()
	},
}

func getLocations() error {
	log.Info().Msg("Retrieving list of locations...")
	
	locations, err := locations.Client.LocationsGet(Location, ProductCategory)

	if err != nil {
		return err
	} else {
		return printer.PrintLocationListResponse(locations)
	}
}
