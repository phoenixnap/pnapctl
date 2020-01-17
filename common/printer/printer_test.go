package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/landoop/tableprinter"
	"phoenixnap.com/pnap-cli/tests/testutil"
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
		name    string
		input   interface{}
		isEmpty bool
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}, false},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}, false},
		{"Empty List", []ExampleStruct1{}, true},
	}

	for _, tc := range testCases {
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
			outputError := MainPrinter.PrintOutput(&tc.input, tc.isEmpty)

			testutil.AssertEqual(test_framework, nil, outputError)
		})
	}
}

func ExamplePrintOutputJsonFormat() {
	printerSetup()

	OutputFormat = "json"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct, false)

	// Output: {
	//     "ID": "123",
	//     "Status": "OK"
	// }
}

func TestPrintOutputYamlFormat(test_framework *testing.T) {
	printerSetup()
	OutputFormat = "yaml"

	testCases := []struct {
		name    string
		input   interface{}
		isEmpty bool
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}, false},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}, false},
		{"Empty List", []ExampleStruct1{}, true},
	}

	for _, tc := range testCases {
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
			outputError := MainPrinter.PrintOutput(tc.input, tc.isEmpty)

			testutil.AssertEqual(test_framework, nil, outputError)
		})
	}
}

func ExamplePrintOutputYamlFormat() {
	printerSetup()

	OutputFormat = "yaml"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	MainPrinter.PrintOutput(inputStruct, false)

	// Output: id: "123"
	// status: OK
}

func TestPrintOutputTableFormat(test_framework *testing.T) {
	printerSetup()
	OutputFormat = ""

	testCases := []struct {
		name     string
		input    interface{}
		isEmpty  bool
		expected string
	}{
		{"Single Element", ExampleStruct1{ID: "123", Status: "OK"}, false, `  ID    STATUS  
 ----- -------- 
  123   OK      
`},
		{"List", []ExampleStruct1{{ID: "123", Status: "OK"}, {ID: "456", Status: "FINE"}}, false, `  ID    STATUS  
 ----- -------- 
  123   OK      
  456   FINE    
`},
		{"Empty", []ExampleStruct1{}, true, ``}, // no output to the table printer. we may still have fmt output
	}

	for _, tc := range testCases {
		test_framework.Run(fmt.Sprintf("%s", tc.name), func(test_framework *testing.T) {
			// Overwriting table printer buffer since it uses a different buffer which we can't check via Example
			customTablePrinterBuffer := new(bytes.Buffer)
			MainPrinter = BodyPrinter{
				Tableprinter: tableprinter.New(customTablePrinterBuffer),
			}

			outputError := MainPrinter.PrintOutput(tc.input, tc.isEmpty)

			testutil.AssertEqual(test_framework, nil, outputError)

			// asserting the custom buffer printed something
			outputText := string(customTablePrinterBuffer.Bytes())

			testutil.AssertEqual(test_framework, tc.expected, outputText)
		})
	}
}

func ExamplePrintOutputTableFormatEmpty() {
	printerSetup()
	OutputFormat = ""

	MainPrinter.PrintOutput([]ExampleStruct1{}, true)

	// Output: No data found
}
