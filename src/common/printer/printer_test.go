package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/landoop/tableprinter"
	"github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/testsupport/testutil"
)

type ExampleStruct1 struct {
	ID     string `header:"id" json:"id"`
	Status string `header:"status" json:"status"`
}

// Setup
func printerSetup() {
	MainPrinter = NewBodyPrinter()
}

func TestPrintOutputJsonFormat(test_framework *testing.T) {
	printerSetup()
	OutputFormat = "json"

	testCases := []struct {
		name  string
		input interface{}
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}},
		{"Empty List", []ExampleStruct1{}},
	}

	for _, tc := range testCases {
		test_framework.Run(tc.name, func(test_framework *testing.T) {
			outputError := MainPrinter.PrintOutput(&tc.input)

			assert.NoError(test_framework, outputError)
		})
	}
}

func ExamplePrintOutputJsonFormat() {
	printerSetup()

	OutputFormat = "json"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct)

	// Output: {
	//     "id": "123",
	//     "status": "OK"
	// }
}

func TestPrintOutputYamlFormat(test_framework *testing.T) {
	printerSetup()
	OutputFormat = "yaml"

	testCases := []struct {
		name  string
		input interface{}
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}},
		{"Empty List", []ExampleStruct1{}},
	}

	for _, tc := range testCases {
		test_framework.Run(tc.name, func(test_framework *testing.T) {
			outputError := MainPrinter.PrintOutput(tc.input)

			assert.NoError(test_framework, outputError)
		})
	}
}

func ExamplePrintOutputYamlFormat() {
	printerSetup()

	OutputFormat = "yaml"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct)

	// Output: id: "123"
	// status: OK
}

func TestPrintOutputTableFormat(test_framework *testing.T) {
	printerSetup()
	OutputFormat = ""

	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}, `  ID    STATUS  
 ----- -------- 
  123   OK      
`},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}, `  ID    STATUS  
 ----- -------- 
  123   OK      
  456   FINE    
`},
		{"Empty", []ExampleStruct1{}, ``}, // no output to the table printer. we may still have fmt output
	}

	for _, tc := range testCases {
		test_framework.Run(tc.name, func(test_framework *testing.T) {
			// Overwriting table printer buffer since it uses a different buffer which we can't check via Example
			customTablePrinterBuffer := new(bytes.Buffer)
			MainPrinter = BodyPrinter{
				Tableprinter: tableprinter.New(customTablePrinterBuffer),
			}

			outputError := MainPrinter.PrintOutput(tc.input)

			assert.NoError(test_framework, outputError)

			// asserting the custom buffer printed something
			outputText := customTablePrinterBuffer.String()

			assert.Equal(test_framework, tc.expected, outputText)
		})
	}
}

func TestPrepareServerForPrintingLongTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.LongServerTable")
}

func TestPrepareServerForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ShortServerTable")
}

func TestPrepareServerForPrintingServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "bmcapi.Server")
}

func TestPrepareServerListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	servers := testutil.GenN(1, generators.Generate[bmcapi.Server])
	prepared := PrepareServerListForPrinting(servers, false)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "bmcapi.Server")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareEventForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "auditapi.Event")
}

func TestPrepareEventForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.Event")
}

func TestPrepareEventListForPrinting(test_framework *testing.T) {
	OutputFormat = "table"
	events := testutil.GenN(1, generators.Generate[auditapi.Event])
	prepared := PrepareEventListForPrinting(events)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.Event")
}

func TestPrepareClusterForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ClusterTable")
}

func TestPrepareClusterForPrintingCluster(test_framework *testing.T) {
	OutputFormat = "json"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "ranchersolutionapi.Cluster")
}

func TestPrepareClusterListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	clusters := testutil.GenN(1, generators.Generate[ranchersolutionapi.Cluster])
	prepared := PrepareClusterListForPrinting(clusters)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "ranchersolutionapi.Cluster")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareTagForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.TagTable")
}

func TestPrepareTagForPrintingTag(test_framework *testing.T) {
	OutputFormat = "json"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tagapi.Tag")
}

func TestPrepareTagListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	tags := testutil.GenN(1, generators.Generate[tagapi.Tag])
	prepared := PrepareTagListForPrinting(tags)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tagapi.Tag")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPreparePrivateNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.PrivateNetworkTable")
}

func TestPreparePrivateNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkapi.PrivateNetwork")
}

func TestPreparePrivateNetworkListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetworks := testutil.GenN(1, generators.Generate[networkapi.PrivateNetwork])
	prepared := PreparePrivateNetworkListForPrinting(privateNetworks)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "networkapi.PrivateNetwork")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareQuotaForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.Quota")
}

func TestPrepareQuotaForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "bmcapi.Quota")
}

func TestPrepareQuotaListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quotas := testutil.GenN(1, generators.Generate[bmcapi.Quota])
	prepared := PrepareQuotaListForPrinting(quotas)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.Quota")
}

func TestPrepareSshkeyFullForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.SshKeyTableFull")
}

func TestPrepareSshkeyForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.SshKeyTable")
}

func TestPrepareSshkeyForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "bmcapi.SshKey")
}

func TestPrepareSshkeyFullListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkeys := testutil.GenN(1, generators.Generate[bmcapi.SshKey])
	prepared := PrepareSshKeyListForPrinting(sshkeys, true)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.SshKeyTableFull")
}

func TestPrepareSshkeyListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkeys := testutil.GenN(1, generators.Generate[bmcapi.SshKey])
	prepared := PrepareSshKeyListForPrinting(sshkeys, false)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.SshKeyTable")
}

func TestPrepareIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.IpBlock")
}

func TestPrepareIpBlockForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.IpBlockShort")
}

func TestPrepareIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "ipapi.IpBlock")
}

func TestPrepareIpBlockListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlocks := testutil.GenN(1, generators.Generate[ipapi.IpBlock])
	prepared := PrepareIpBlockListForPrinting(ipBlocks, true)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.IpBlock")
}

func TestPrepareIpBlockListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlocks := testutil.GenN(1, generators.Generate[ipapi.IpBlock])
	prepared := PrepareIpBlockListForPrinting(ipBlocks, true)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "ipapi.IpBlock")
	assert.Equal(test_framework, len(prepared), 1)
}

// Billing

func TestPrepareRatedUsageRecordForPrintingNonTable_Bandwidth(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateBandwidthRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.BandwidthRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_OperatingSystem(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateOperatingSystemRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.OperatingSystemRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_PublicSubnet(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GeneratePublicSubnetRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.PublicSubnetRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Server(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.ServerRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Storage(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateStorageRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.StorageRecord")
}

func TestPrepareRatedUsageRecordForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*tables.RatedUsageRecordTable")
}

func TestPrepareRatedUsageRecordForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*tables.ShortRatedUsageRecordTable")
}

func TestPrepareProductForPrintingNonTable_BandwidthProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.Product")
}

func TestPrepareProductForPrintingNonTable_OperatingSystemProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateOperatingSystemProduct()
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.Product")
}

func TestPrepareProductForPrintingNonTable_StorageProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateStorageProduct()
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.Product")
}

func TestPrepareProductForPrintingNonTable_ServerProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateServerProduct()
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingapi.ServerProduct")
}

func TestPrepareProductForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*tables.ProductTable")
}

func TestPreparePublicNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.PublicNetworkTable")
}

func TestPreparePublicNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkapi.PublicNetwork")
}

func TestPreparePublicNetworkIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.PublicNetworkIpBlockTable")
}

func TestPreparePublicNetworkIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkapi.PublicNetworkIpBlock")
}

func TestPrepareStorageNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.StorageNetworkTable")
}

func TestPrepareStorageNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkstorageapi.StorageNetwork")
}

func TestPrepareVolumeForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.VolumeTable")
}

func TestPrepareVolumeForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ShortVolumeTable")
}

func TestPrepareVolumeForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkstorageapi.Volume")
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{})

	// Output: No data found.
}
