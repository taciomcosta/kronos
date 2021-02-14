package main

import (
	"fmt"
	"os"

	"github.com/taciomcosta/kronos/internal/interfaces/cli"
)

func main() {
	client := cli.NewClient()
	err := client.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
