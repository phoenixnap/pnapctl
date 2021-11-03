package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/landoop/tableprinter"
	"phoenixnap.com/pnap-cli/tests/generators"
)

type ExampleStruct1 struct {
	ID     string `header:"id"`
	Status string `header:"status"`
}

var input []byte

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
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
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
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
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
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
			// Overwriting table printer buffer since it uses a different buffer which we can't check via Example
			customTablePrinterBuffer := new(bytes.Buffer)
			MainPrinter = BodyPrinter{
				Tableprinter: tableprinter.New(customTablePrinterBuffer),
			}

			outputError := MainPrinter.PrintOutput(tc.input, "dummy command")

			assert.NoError(test_framework, outputError)

			// asserting the custom buffer printed something
			outputText := string(customTablePrinterBuffer.Bytes())

			assert.Equal(test_framework, tc.expected, outputText)
		})
	}
}

func TestPrepareServerForPrintingLongTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.GenerateServer()
	prepared := PrepareServerForPrinting(server, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.LongServerTable")
}

func TestPrepareServerForPrintingShortTable(test_framework *testing.T) {
	OutputFormat = "table"
	server := generators.GenerateServer()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "tables.ShortServerTable")
}

func TestPrepareServerForPrintingLongServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := generators.GenerateServer()
	prepared := PrepareServerForPrinting(server, true)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "models.LongServer")
}

func TestPrepareServerForPrintingShortServer(test_framework *testing.T) {
	OutputFormat = "json"
	server := generators.GenerateServer()
	prepared := PrepareServerForPrinting(server, false)

	outputType := fmt.Sprintf("%T", prepared)

	assert.Equal(test_framework, outputType, "models.ShortServer")
}

func TestPrepareServerListForPrinting(test_framework *testing.T) {
	OutputFormat = "json"
	servers := generators.GenerateServers(1)
	prepared := PrepareServerListForPrinting(servers, false)

	outputType := fmt.Sprintf("%T", prepared[0])

	assert.Equal(test_framework, outputType, "models.ShortServer")
	assert.Equal(test_framework, len(prepared), 1)
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{}, "dummy command")

	// Output: No data found.
}
