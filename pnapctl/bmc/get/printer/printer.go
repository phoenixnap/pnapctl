package printer

import "fmt"

var OutputFormat string

func PrintOutput(body []byte, construct interface{}) {
	fmt.Println("Printing body with the format:", OutputFormat)
}
