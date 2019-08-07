package printer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
)

var tblprinter = tableprinter.New(os.Stdout)

var OutputFormat string

// PrintOutput prints the construct passed according to the format.
// The first parameter is only used by the table printer to show how
// many rows were printed. The second parameter specifies any errors.
func PrintOutput(body []byte, construct interface{}) (int, error) {
	fmt.Println("Printing body with the format:", OutputFormat)

	err := json.Unmarshal(body, &construct)

	if err != nil {
		return 0, err
	}

	if OutputFormat == "json" {
		printJSON(body)
		return -1, nil
	} else if OutputFormat == "yaml" {
		err := printYAML(construct)
		return -1, err
	} else {
		// default to table
		rows := printTable(construct)
		if rows == -1 {
			return -1, errors.New("table-print-failed")
		}
		return rows, nil
	}
}

func printYAML(body interface{}) error {
	b, err := yaml.Marshal(body)
	fmt.Println(string(b))
	return err
}

func printJSON(body []byte) {
	var dat bytes.Buffer
	json.Indent(&dat, body, "", "    ")
	fmt.Println(string(dat.Bytes()))
}

func printTable(body interface{}) int {
	return tblprinter.Print(body)
}
