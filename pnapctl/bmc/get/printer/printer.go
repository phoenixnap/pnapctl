package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/landoop/tableprinter"
	"gopkg.in/yaml.v2"
)

var tblprinter = tableprinter.New(os.Stdout)

var OutputFormat string

func PrintOutput(body []byte, construct interface{}) {
	fmt.Println("Printing body with the format:", OutputFormat)

	json.Unmarshal(body, &construct)

	if OutputFormat == "json" {
		printJSON(body)
	} else if OutputFormat == "yaml" {
		printYAML(construct)
	} else {
		// default to table
		printTable(construct)
	}
}

func printYAML(body interface{}) {
	b, _ := yaml.Marshal(body)
	fmt.Println(string(b))
}

func printJSON(body []byte) {
	var dat bytes.Buffer
	json.Indent(&dat, body, "", "    ")
	fmt.Println(string(dat.Bytes()))
}

func printTable(body interface{}) {
	tblprinter.Print(body)
}
