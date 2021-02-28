package cli

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [jobs | channels]",
	Short: "List jobs/channels",
}

func init() {
	listCmd.AddCommand(listJobsCmd)
}
