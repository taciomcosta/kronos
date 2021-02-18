package main

import (
	"fmt"
	"os"

	"github.com/taciomcosta/kronos/internal/config"
	"github.com/taciomcosta/kronos/internal/interfaces/cli"
)

func main() {
	_ = config.EnableDefaultMode()

	kronosdURL := "http://localhost" + config.GetString("host")
	client := cli.NewClient(kronosdURL)

	err := client.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
