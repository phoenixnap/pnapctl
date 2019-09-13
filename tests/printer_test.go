package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	"github.com/landoop/tableprinter"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
	"phoenixnap.com/pnap-cli/pnapctl/printer"
	"phoenixnap.com/pnap-cli/tests/testutil"
)

type ExampleStruct1 struct {
	ID     string `header:"id"`
	Status string `header:"status"`
}

var input []byte

// Setup
func TestMain(m *testing.M) {
	printer.MainPrinter = printer.NewBodyPrinter()
}

func TestPrintOutputUnmarshalError(test_framework *testing.T) {
	input := []byte("test input")

	errPO := printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	expectedErr := errors.New(ctlerrors.UnmarshallingInPrinter)

	testutil.AssertEqual(test_framework, expectedErr.Error(), errPO.Error())
}

func ExamplePrintOutputUnmarshalError() {
	input := []byte("test input")

	printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	// Output:
}

func TestPrintOutputJsonFormat(test_framework *testing.T) {
	printer.OutputFormat = "json"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	input, _ = json.Marshal(inputStruct)

	outputError := printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	testutil.AssertEqual(test_framework, nil, outputError)
}

func ExamplePrintOutputJsonFormat() {
	printer.OutputFormat = "json"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	input, _ = json.Marshal(inputStruct)

	printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	// Output: {
	//     "ID": "123",
	//     "Status": "OK"
	// }
}

func TestPrintOutputYamlFormat(test_framework *testing.T) {
	printer.OutputFormat = "yaml"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	input, _ = json.Marshal(inputStruct)

	outputError := printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	testutil.AssertEqual(test_framework, nil, outputError)
}

func ExamplePrintOutputYamlFormat() {
	printer.OutputFormat = "yaml"

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	input, _ = json.Marshal(inputStruct)

	printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	// Output: id: "123"
	// status: OK
}

func TestPrintOutputTableFormat(test_framework *testing.T) {
	printer.OutputFormat = ""

	// Overwriting table printer buffer since it uses a different buffer which we can't check via Example
	customTablePrinterBuffer := new(bytes.Buffer)
	printer.MainPrinter = printer.BodyPrinter{
		Tableprinter: tableprinter.New(customTablePrinterBuffer),
	}

	inputStruct := ExampleStruct1{ID: "123", Status: "OK"}

	input, _ = json.Marshal(inputStruct)

	outputError := printer.MainPrinter.PrintOutput(input, &ExampleStruct1{})

	testutil.AssertEqual(test_framework, nil, outputError)

	// asserting the custom buffer printed something
	outputText := string(customTablePrinterBuffer.Bytes())
	outputTextLength := len(outputText)

	if outputTextLength == 0 {
		test_framework.Error("Table printer did not print anything")
	}
}
