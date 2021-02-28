package cli

import (
	"github.com/spf13/cobra"
)

var kronosdURL string
var flags = struct {
	Name    string
	Command string
	Tick    string
}{}

// NewClient creates a new CLI client
func NewClient(url string) *cobra.Command {
	kronosdURL = url
	return rootCmd
}
