package printer

import (
	"bytes"
	"fmt"
	"testing"

	"phoenixnap.com/pnapctl/common/models/billingmodels"
	"phoenixnap.com/pnapctl/common/models/ipmodels"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/landoop/tableprinter"
	"github.com/phoenixnap/go-sdk-bmc/billingapi"
	"phoenixnap.com/pnapctl/common/models/auditmodels"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/quotamodels"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/servermodels"
	"phoenixnap.com/pnapctl/common/models/bmcapimodels/sshkeymodels"
	"phoenixnap.com/pnapctl/common/models/networkmodels"
	"phoenixnap.com/pnapctl/common/models/ranchermodels"
	"phoenixnap.com/pnapctl/common/models/tagmodels"
)

type ExampleStruct1 struct {
	ID     string `header:"id"`
	Status string `header:"status"`
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
			outputError := MainPrinter.PrintOutput(&tc.input, "dummy command")

			assert.NoError(test_framework, outputError)
		})
	}
}

func ExamplePrintOutputJsonFormat() {
	printerSetup()

	OutputFormat = "json"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct, "dummy command")

	// Output: {
	//     "ID": "123",
	//     "Status": "OK"
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
			outputError := MainPrinter.PrintOutput(tc.input, "dummy command")

			assert.NoError(test_framework, outputError)
		})
	}
}

func ExamplePrintOutputYamlFormat() {
	printerSetup()

	OutputFormat = "yaml"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct, "dummy command")

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

			outputError := MainPrinter.PrintOutput(tc.input, "dummy command")

			assert.NoError(test_framework, outputError)

			// asserting the custom buffer printed something
			outputText := customTablePrinterBuffer.String()

			assert.Equal(test_framework, tc.expected, outputText)
		})
	}
}

func TestPrepareServerForPrintingLongTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := servermodels.GenerateServerSdk()
	prepared := PrepareServerForPrinting(server, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.LongServerTable")
}

func TestPrepareServerForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := servermodels.GenerateServerSdk()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ShortServerTable")
}

func TestPrepareServerForPrintingLongServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := servermodels.GenerateServerSdk()
	prepared := PrepareServerForPrinting(server, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "servermodels.LongServer")
}

func TestPrepareServerForPrintingShortServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := servermodels.GenerateServerSdk()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "servermodels.ShortServer")
}

func TestPrepareServerListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	servers := servermodels.GenerateServerListSdk(1)
	prepared := PrepareServerListForPrinting(servers, false)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "servermodels.ShortServer")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareEventForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	event := auditmodels.GenerateEventSdk()
	prepared := PrepareEventForPrinting(*event)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*auditmodels.Event")
}

func TestPrepareEventForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	event := auditmodels.GenerateEventSdk()
	prepared := PrepareEventForPrinting(*event)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.Event")
}

func TestPrepareEventListForPrinting(test_framework *testing.T) {
	OutputFormat = "table"
	events := auditmodels.GenerateEventListSdk(1)
	prepared := PrepareEventListForPrinting(events)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.Event")
}

func TestPrepareClusterForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	cluster := ranchermodels.GenerateClusterSdk()
	prepared := PrepareClusterForPrinting(cluster)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ClusterTable")
}

func TestPrepareClusterForPrintingCluster(test_framework *testing.T) {
	OutputFormat = "json"
	cluster := ranchermodels.GenerateClusterSdk()
	prepared := PrepareClusterForPrinting(cluster)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "ranchermodels.Cluster")
}

func TestPrepareClusterListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	clusters := ranchermodels.GenerateClusterListSdk(1)
	prepared := PrepareClusterListForPrinting(clusters)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "ranchermodels.Cluster")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareTagForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	tag := tagmodels.GenerateTagSdk()
	prepared := PrepareTagForPrinting(*tag)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.TagTable")
}

func TestPrepareTagForPrintingTag(test_framework *testing.T) {
	OutputFormat = "json"
	tag := *tagmodels.GenerateTagSdk()
	prepared := PrepareTagForPrinting(tag)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tagmodels.Tag")
}

func TestPrepareTagListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	tags := tagmodels.GenerateTagListSdk(1)
	prepared := PrepareTagListForPrinting(tags)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tagmodels.Tag")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPreparePrivateNetworkForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	privateNetwork := networkmodels.GeneratePrivateNetworkSdk()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.PrivateNetworkTable")
}

func TestPreparePrivateNetworkForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetwork := networkmodels.GeneratePrivateNetworkSdk()
	prepared := PreparePrivateNetworkForPrinting(privateNetwork)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "networkmodels.PrivateNetwork")
}

func TestPreparePrivateNetworkListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	privateNetworks := networkmodels.GeneratePrivateNetworkListSdk(1)
	prepared := PreparePrivateNetworkListForPrinting(privateNetworks)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "networkmodels.PrivateNetwork")
	assert.Equal(test_framework, len(prepared), 1)
}

func TestPrepareQuotaForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quota := quotamodels.GenerateQuotaSdk()
	prepared := PrepareQuotaForPrinting(quota)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.Quota")
}

func TestPrepareQuotaForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	quota := quotamodels.GenerateQuotaSdk()
	prepared := PrepareQuotaForPrinting(quota)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "quotamodels.Quota")
}

func TestPrepareQuotaListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	quotas := quotamodels.GenerateQuotaSdkList(1)
	prepared := PrepareQuotaListForPrinting(quotas)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.Quota")
}

func TestPrepareSshkeyFullForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := sshkeymodels.GenerateSshKeySdk()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.SshKeyTableFull")
}

func TestPrepareSshkeyForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkey := sshkeymodels.GenerateSshKeySdk()
	prepared := PrepareSshKeyForPrinting(sshkey, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.SshKeyTable")
}

func TestPrepareSshkeyForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	sshkey := sshkeymodels.GenerateSshKeySdk()
	prepared := PrepareSshKeyForPrinting(sshkey, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "sshkeymodels.SshKey")
}

func TestPrepareSshkeyFullListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkeys := sshkeymodels.GenerateSshKeyListSdk(1)
	prepared := PrepareSshKeyListForPrinting(sshkeys, true)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.SshKeyTableFull")
}

func TestPrepareSshkeyListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	sshkeys := sshkeymodels.GenerateSshKeyListSdk(1)
	prepared := PrepareSshKeyListForPrinting(sshkeys, false)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.SshKeyTable")
}

func TestPrepareIpBlockForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlock := ipmodels.GenerateIpBlockSdk()
	prepared := PrepareIpBlockForPrinting(ipBlock)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.IpBlock")
}

func TestPrepareIpBlockForPrintingNonTable(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlock := ipmodels.GenerateIpBlockSdk()
	prepared := PrepareIpBlockForPrinting(ipBlock)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "ipmodels.IpBlock")
}

func TestPrepareIpBlockListForPrintingTable(test_framework *testing.T) {
	OutputFormat = "table"
	ipBlocks := ipmodels.GenerateIpBlockSdkList(1)
	prepared := PrepareIpBlockListForPrinting(ipBlocks)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "tables.IpBlock")
}

func TestPrepareIpBlockListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	ipBlocks := ipmodels.GenerateIpBlockSdkList(1)
	prepared := PrepareIpBlockListForPrinting(ipBlocks)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "ipmodels.IpBlock")
	assert.Equal(test_framework, len(prepared), 1)
}

// Billing

func TestPrepareRatedUsageRecordForPrintingNonTable_Bandwidth(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: billingmodels.GenerateBandwidthRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*ratedusageoneof.BandwidthRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_OperatingSystem(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		OperatingSystemRecord: billingmodels.GenerateOperatingSystemRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*ratedusageoneof.OperatingSystemRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_PublicSubnet(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		PublicSubnetRecord: billingmodels.GeneratePublicSubnetRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*ratedusageoneof.PublicSubnetRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Server(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: billingmodels.GenerateServerRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*ratedusageoneof.ServerRecord")
}

func TestPrepareRatedUsageRecordForPrintingNonTable_Short(test_framework *testing.T) {
	OutputFormat = "json"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		BandwidthRecord: billingmodels.GenerateBandwidthRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*billingmodels.ShortRatedUsage")
}

func TestPrepareRatedUsageRecordForPrintingTableFull(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: billingmodels.GenerateServerRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.RatedUsageRecordTable")
}

func TestPrepareRatedUsageRecordForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	ratedUsage := billingapi.RatedUsageGet200ResponseInner{
		ServerRecord: billingmodels.GenerateServerRecordSdk(),
	}
	prepared := PrepareRatedUsageForPrinting(ratedUsage, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ShortRatedUsageRecordTable")
}

func TestPrepareProductForPrintingNonTable_BandwidthProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := billingapi.ProductsGet200ResponseInner{
		Product: billingmodels.GenerateBandwidthProduct(),
	}
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*productoneof.Product")
}

func TestPrepareProductForPrintingNonTable_OperatingSystemProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := billingapi.ProductsGet200ResponseInner{
		Product: billingmodels.GenerateOperatingSystemProduct(),
	}
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*productoneof.Product")
}

func TestPrepareProductForPrintingNonTable_ServerProduct(test_framework *testing.T) {
	OutputFormat = "json"
	product := billingapi.ProductsGet200ResponseInner{
		ServerProduct: billingmodels.GenerateServerProduct(),
	}
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "*productoneof.ServerProduct")
}

func TestPrepareProductForPrintingTableShort(test_framework *testing.T) {
	OutputFormat = "table"
	product := billingapi.ProductsGet200ResponseInner{
		Product: billingmodels.GenerateBandwidthProduct(),
	}
	prepared := PrepareProductForPrinting(product)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ProductTable")
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{}, "dummy command")

	// Output: No data found.
}
