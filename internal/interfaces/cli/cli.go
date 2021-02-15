package cli

import (
	"github.com/spf13/cobra"
)

var url = "http://localhost:8080"

var flags = struct {
	Name    string
	Command string
	Tick    string
}{}

func setup() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)

	createCmd.AddCommand(createJobCmd)

	createJobCmd.Flags().StringVarP(&flags.Name, "name", "n", "", "Unique job name")
	createJobCmd.Flags().StringVarP(&flags.Command, "cmd", "c", "", "Job entrypoint")
	createJobCmd.Flags().StringVarP(&flags.Tick, "tick", "t", "", `Cron expression. Ex: "* * * * *"`)
	_ = createJobCmd.MarkFlagRequired("name")
	_ = createJobCmd.MarkFlagRequired("cmd")
	_ = createJobCmd.MarkFlagRequired("tick")

	listCmd.AddCommand(listJobsCmd)
}

// NewClient creates a new CLI client
func NewClient() *cobra.Command {
	setup()
	return rootCmd
}
