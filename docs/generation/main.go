package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	root "phoenixnap.com/pnapctl/apps/pnapctl"
)

func main() {
	root.RootCmd.DisableAutoGenTag = true
	err := doc.GenMarkdownTree(root.RootCmd, "../")
	if err != nil {
		log.Fatal(err)
	}
}
