package printer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
	"phoenixnap.com/pnap-cli/pnapctl/ctlerrors"
)

// The main printer used by the application.
// Can be swapped by a Mock for testing.
var MainPrinter = NewBodyPrinter()

// The format to print the object in.
var OutputFormat string

type Printer interface {
	PrintOutput(body []byte, construct interface{}) (int, error)
}

type BodyPrinter struct {
	tableprinter *tableprinter.Printer
}

func NewBodyPrinter() Printer {
	return BodyPrinter{
		tableprinter: tableprinter.New(os.Stdout),
	}
}

// PrintOutput prints the construct passed according to the format.
// The first output parameter is only used by the table printer to show how
// many rows were printed. The second output parameter specifies any errors.
func (m BodyPrinter) PrintOutput(body []byte, construct interface{}) (int, error) {
	err := json.Unmarshal(body, &construct)

	if err != nil {
		return -1, errors.New(ctlerrors.UnmarshallingInPrinter)
	}

	if OutputFormat == "json" {
		printJSON(body)
		return 0, nil
	} else if OutputFormat == "yaml" {
		err := printYAML(construct)
		return -1, err
	} else {
		// default to table
		rows := printTable(construct, m.tableprinter)
		if rows == -1 {
			return -1, errors.New(ctlerrors.TablePrinterFailure)
		}
		return rows, nil
	}
}

// Attempts to print in YAML via marshalling.
func printYAML(body interface{}) error {
	b, err := yaml.Marshal(body)
	fmt.Println(string(b))

	if err != nil {
		return errors.New(ctlerrors.MarshallingInPrinter)
	}

	return nil
}

// Attempts to print in JSON via formatting a byte array.
func printJSON(body []byte) {
	var dat bytes.Buffer
	json.Indent(&dat, body, "", "    ")
	fmt.Println(string(dat.Bytes()))
}

// Attempts to print the struct as a table.
func printTable(body interface{}, tblprinter *tableprinter.Printer) int {
	return tblprinter.Print(body)
}
