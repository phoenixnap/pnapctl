package get

import (
	"os"

	"github.com/spf13/cobra"
	account_billing_configuration "phoenixnap.com/pnapctl/commands/get/account-billing-configuration"
	"phoenixnap.com/pnapctl/commands/get/bgppeergroup"
	"phoenixnap.com/pnapctl/commands/get/clusters"
	"phoenixnap.com/pnapctl/commands/get/events"
	"phoenixnap.com/pnapctl/commands/get/invoicing"
	ip_blocks "phoenixnap.com/pnapctl/commands/get/ip-blocks"
	"phoenixnap.com/pnapctl/commands/get/locations"
	"phoenixnap.com/pnapctl/commands/get/privatenetwork"
	product_availability "phoenixnap.com/pnapctl/commands/get/product-availability"
	"phoenixnap.com/pnapctl/commands/get/products"
	"phoenixnap.com/pnapctl/commands/get/publicnetwork"
	"phoenixnap.com/pnapctl/commands/get/quotas"
	rated_usage "phoenixnap.com/pnapctl/commands/get/rated-usage"
	"phoenixnap.com/pnapctl/commands/get/reservations"
	"phoenixnap.com/pnapctl/commands/get/servers"
	"phoenixnap.com/pnapctl/commands/get/sshkeys"
	storagenetworks "phoenixnap.com/pnapctl/commands/get/storage-networks"
	"phoenixnap.com/pnapctl/commands/get/tags"
	"phoenixnap.com/pnapctl/commands/get/transactions"
)

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources.",
	Long:  `Display one or many resources.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	GetCmd.AddCommand(servers.GetServersCmd)
	GetCmd.AddCommand(clusters.GetClustersCmd)
	GetCmd.AddCommand(quotas.GetQuotasCmd)
	GetCmd.AddCommand(events.GetEventsCmd)
	GetCmd.AddCommand(tags.GetTagsCmd)
	GetCmd.AddCommand(sshkeys.GetSshKeysCmd)
	GetCmd.AddCommand(privatenetwork.GetPrivateNetworksCmd)
	GetCmd.AddCommand(publicnetwork.GetPublicNetworksCmd)
	GetCmd.AddCommand(bgppeergroup.GetBgpPeerGroupsCmd)
	GetCmd.AddCommand(ip_blocks.GetIpBlockCmd)
	GetCmd.AddCommand(rated_usage.GetRatedUsageCmd)
	GetCmd.AddCommand(products.GetProductsCmd)
	GetCmd.AddCommand(reservations.GetReservationsCmd)
	GetCmd.AddCommand(product_availability.GetProductAvailabilitiesCmd)
	GetCmd.AddCommand(account_billing_configuration.GetAccountBillingConfigurationCmd)
	GetCmd.AddCommand(storagenetworks.GetStorageNetworksCmd)
	GetCmd.AddCommand(locations.GetLocationsCmd)
	GetCmd.AddCommand(invoicing.GetInvoicingCmd)
	GetCmd.AddCommand(transactions.GetTransactionsCmd)
}
