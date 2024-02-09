package printer

import (
	"bytes"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/landoop/tableprinter"
	"github.com/phoenixnap/go-sdk-bmc/auditapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/billingapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/locationapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v3"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
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

	testutil.AssertIsType[tables.LongServerTable](test_framework, prepared)
}

func TestPrepareServerForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	testutil.AssertIsType[tables.ShortServerTable](test_framework, prepared)
}

func TestPrepareServerForPrintingServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	testutil.AssertIsType[bmcapi.Server](test_framework, prepared)
}

func TestPrepareEventForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	testutil.AssertIsType[auditapi.Event](test_framework, prepared)
}

func TestPrepareEventForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	testutil.AssertIsType[tables.Event](test_framework, prepared)
}

func TestPrepareClusterForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	testutil.AssertIsType[tables.ClusterTable](test_framework, prepared)
}

func TestPrepareClusterForPrintingCluster(test_framework *testing.T) {
	OutputFormat = "json"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	testutil.AssertIsType[ranchersolutionapi.Cluster](test_framework, prepared)
}

func TestPrepareTagForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	testutil.AssertIsType[tables.TagTable](test_framework, prepared)
}

func TestPrepareTagForPrintingTag(test_framework *testing.T) {
	OutputFormat = "json"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	testutil.AssertIsType[tagapi.Tag](test_framework, prepared)
}

func TestPreparePrivateNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	testutil.AssertIsType[tables.PrivateNetworkTable](test_framework, prepared)
}

func TestPreparePrivateNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	testutil.AssertIsType[networkapi.PrivateNetwork](test_framework, prepared)
}

func TestPrepareQuotaForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	testutil.AssertIsType[tables.Quota](test_framework, prepared)
}

func TestPrepareQuotaForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	testutil.AssertIsType[bmcapi.Quota](test_framework, prepared)
}

func TestPrepareSshkeyFullForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	testutil.AssertIsType[tables.SshKeyTableFull](test_framework, prepared)
}

func TestPrepareSshkeyForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, false)

	testutil.AssertIsType[tables.SshKeyTable](test_framework, prepared)
}

func TestPrepareSshkeyForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	testutil.AssertIsType[bmcapi.SshKey](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	testutil.AssertIsType[tables.IpBlock](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, false)

	testutil.AssertIsType[tables.IpBlockShort](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	testutil.AssertIsType[ipapi.IpBlock](test_framework, prepared)
}

// Billing

func TestPrepareRatedUsageRecordForPrintingNonTable_Bandwidth(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateBandwidthRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*billingapi.BandwidthRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_OperatingSystem(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateOperatingSystemRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*billingapi.OperatingSystemRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_PublicSubnet(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GeneratePublicSubnetRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*billingapi.PublicSubnetRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Server(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*billingapi.ServerRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Storage(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateStorageRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*billingapi.StorageRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	testutil.AssertIsType[*tables.RatedUsageRecordTable](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, false)

	testutil.AssertIsType[*tables.ShortRatedUsageRecordTable](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_BandwidthProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	testutil.AssertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_OperatingSystemProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateOperatingSystemProduct()
	prepared := PrepareProductForPrinting(product)

	testutil.AssertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_StorageProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateStorageProduct()
	prepared := PrepareProductForPrinting(product)

	testutil.AssertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_ServerProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateServerProduct()
	prepared := PrepareProductForPrinting(product)

	testutil.AssertIsType[*billingapi.ServerProduct](test_framework, prepared)
}

func TestPrepareProductForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	testutil.AssertIsType[*tables.ProductTable](test_framework, prepared)
}

func TestPreparePublicNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	testutil.AssertIsType[tables.PublicNetworkTable](test_framework, prepared)
}

func TestPreparePublicNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	testutil.AssertIsType[networkapi.PublicNetwork](test_framework, prepared)
}

func TestPreparePublicNetworkIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	testutil.AssertIsType[tables.PublicNetworkIpBlockTable](test_framework, prepared)
}

func TestPreparePublicNetworkIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	testutil.AssertIsType[networkapi.PublicNetworkIpBlock](test_framework, prepared)
}

func TestPrepareStorageNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	testutil.AssertIsType[tables.StorageNetworkTable](test_framework, prepared)
}

func TestPrepareStorageNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	testutil.AssertIsType[networkstorageapi.StorageNetwork](test_framework, prepared)
}

func TestPrepareVolumeForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	testutil.AssertIsType[tables.VolumeTable](test_framework, prepared)
}

func TestPrepareVolumeForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, false)

	testutil.AssertIsType[tables.ShortVolumeTable](test_framework, prepared)
}

func TestPrepareVolumeForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	testutil.AssertIsType[networkstorageapi.Volume](test_framework, prepared)
}

func TestPrepareLocationForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	location := generators.Generate[locationapi.Location]()
	prepared := PrepareLocationForPrinting(location)

	testutil.AssertIsType[tables.Location](test_framework, prepared)
}

func TestPrepareLocationForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	location := generators.Generate[locationapi.Location]()
	prepared := PrepareLocationForPrinting(location)

	testutil.AssertIsType[locationapi.Location](test_framework, prepared)
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{})

	// Output: No data found.
}
