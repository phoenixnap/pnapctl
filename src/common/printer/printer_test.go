package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/landoop/tableprinter"
	"github.com/phoenixnap/go-sdk-bmc/auditapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"github.com/phoenixnap/go-sdk-bmc/bmcapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/ipapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/networkapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/networkstorageapi"
	"github.com/phoenixnap/go-sdk-bmc/ranchersolutionapi/v2"
	"github.com/phoenixnap/go-sdk-bmc/tagapi/v2"
	"phoenixnap.com/pnapctl/common/models/generators"
	"phoenixnap.com/pnapctl/common/models/tables"
)

type ExampleStruct1 struct {
	ID     string `header:"id" json:"id"`
	Status string `header:"status" json:"status"`
}

// Setup
func printerSetup() {
	MainPrinter = NewBodyPrinter()
}

func assertIsType[T any](test_framework *testing.T, item interface{}) {
	var t T
	assert.Equal(
		test_framework,
		fmt.Sprintf("%T", item),
		fmt.Sprintf("%T", t),
	)
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

	assertIsType[tables.LongServerTable](test_framework, prepared)
}

func TestPrepareServerForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	assertIsType[tables.ShortServerTable](test_framework, prepared)
}

func TestPrepareServerForPrintingServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := generators.Generate[bmcapi.Server]()
	prepared := PrepareServerForPrinting(server, false)

	assertIsType[bmcapi.Server](test_framework, prepared)
}

func TestPrepareEventForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	assertIsType[auditapi.Event](test_framework, prepared)
}

func TestPrepareEventForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	event := generators.Generate[auditapi.Event]()
	prepared := PrepareEventForPrinting(event)

	assertIsType[tables.Event](test_framework, prepared)
}

func TestPrepareClusterForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	assertIsType[tables.ClusterTable](test_framework, prepared)
}

func TestPrepareClusterForPrintingCluster(test_framework *testing.T) {
	OutputFormat = "json"
	cluster := generators.Generate[ranchersolutionapi.Cluster]()
	prepared := PrepareClusterForPrinting(cluster)

	assertIsType[ranchersolutionapi.Cluster](test_framework, prepared)
}

func TestPrepareTagForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	assertIsType[tables.TagTable](test_framework, prepared)
}

func TestPrepareTagForPrintingTag(test_framework *testing.T) {
	OutputFormat = "json"
	tag := generators.Generate[tagapi.Tag]()
	prepared := PrepareTagForPrinting(tag)

	assertIsType[tagapi.Tag](test_framework, prepared)
}

func TestPreparePrivateNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	assertIsType[tables.PrivateNetworkTable](test_framework, prepared)
}

func TestPreparePrivateNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetwork := generators.Generate[networkapi.PrivateNetwork]()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	assertIsType[networkapi.PrivateNetwork](test_framework, prepared)
}

func TestPrepareQuotaForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	assertIsType[tables.Quota](test_framework, prepared)
}

func TestPrepareQuotaForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	quota := generators.Generate[bmcapi.Quota]()
	prepared := PrepareQuotaForPrinting(quota)

	assertIsType[bmcapi.Quota](test_framework, prepared)
}

func TestPrepareSshkeyFullForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	assertIsType[tables.SshKeyTableFull](test_framework, prepared)
}

func TestPrepareSshkeyForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, false)

	assertIsType[tables.SshKeyTable](test_framework, prepared)
}

func TestPrepareSshkeyForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	sshkey := generators.Generate[bmcapi.SshKey]()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	assertIsType[bmcapi.SshKey](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	assertIsType[tables.IpBlock](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, false)

	assertIsType[tables.IpBlockShort](test_framework, prepared)
}

func TestPrepareIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[ipapi.IpBlock]()
	prepared := PrepareIpBlockForPrinting(ipBlock, true)

	assertIsType[ipapi.IpBlock](test_framework, prepared)
}

// Billing

func TestPrepareRatedUsageRecordForPrintingNonTable_Bandwidth(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateBandwidthRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*billingapi.BandwidthRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_OperatingSystem(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateOperatingSystemRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*billingapi.OperatingSystemRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_PublicSubnet(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GeneratePublicSubnetRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*billingapi.PublicSubnetRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Server(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*billingapi.ServerRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Storage(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := generators.GenerateStorageRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*billingapi.StorageRecord](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	assertIsType[*tables.RatedUsageRecordTable](test_framework, prepared)
}

func TestPrepareRatedUsageRecordForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := generators.GenerateServerRecordSdk()
	prepared := PrepareRatedUsageForPrinting(ratedUsage, false)

	assertIsType[*tables.ShortRatedUsageRecordTable](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_BandwidthProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	assertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_OperatingSystemProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateOperatingSystemProduct()
	prepared := PrepareProductForPrinting(product)

	assertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_StorageProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateStorageProduct()
	prepared := PrepareProductForPrinting(product)

	assertIsType[*billingapi.Product](test_framework, prepared)
}

func TestPrepareProductForPrintingNonTable_ServerProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := generators.GenerateServerProduct()
	prepared := PrepareProductForPrinting(product)

	assertIsType[*billingapi.ServerProduct](test_framework, prepared)
}

func TestPrepareProductForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	product := generators.GenerateBandwidthProduct()
	prepared := PrepareProductForPrinting(product)

	assertIsType[*tables.ProductTable](test_framework, prepared)
}

func TestPreparePublicNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	assertIsType[tables.PublicNetworkTable](test_framework, prepared)
}

func TestPreparePublicNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	publicNetwork := generators.Generate[networkapi.PublicNetwork]()
	prepared := PreparePublicNetworkForPrinting(publicNetwork)

	assertIsType[networkapi.PublicNetwork](test_framework, prepared)
}

func TestPreparePublicNetworkIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	assertIsType[tables.PublicNetworkIpBlockTable](test_framework, prepared)
}

func TestPreparePublicNetworkIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := generators.Generate[networkapi.PublicNetworkIpBlock]()
	prepared := PreparePublicNetworkIpBlockForPrinting(ipBlock)

	assertIsType[networkapi.PublicNetworkIpBlock](test_framework, prepared)
}

func TestPrepareStorageNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	assertIsType[tables.StorageNetworkTable](test_framework, prepared)
}

func TestPrepareStorageNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.StorageNetwork]()
	prepared := PrepareNetworkStorageForPrinting(networkStorage)

	assertIsType[networkstorageapi.StorageNetwork](test_framework, prepared)
}

func TestPrepareVolumeForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	assertIsType[tables.VolumeTable](test_framework, prepared)
}

func TestPrepareVolumeForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, false)

	assertIsType[tables.ShortVolumeTable](test_framework, prepared)
}

func TestPrepareVolumeForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	networkStorage := generators.Generate[networkstorageapi.Volume]()
	prepared := PrepareVolumeForPrinting(networkStorage, true)

	assertIsType[networkstorageapi.Volume](test_framework, prepared)
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{})

	// Output: No data found.
}
