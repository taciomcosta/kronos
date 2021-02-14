package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taciomcosta/kronos/internal/usecases"
)

var rootCmd = &cobra.Command{
	Use:   "kronos",
	Short: "kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs",
	Long:  "kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs",
}

var createCmd = &cobra.Command{
	Use:   "create [job | channel]",
	Short: "Creates a new job/channel",
}

var createJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Creates a new job",
	Run: func(cmd *cobra.Command, args []string) {
		createJobRequest := usecases.CreateJobRequest{
			Name:    flags.Name,
			Command: flags.Command,
			Tick:    flags.Tick,
		}
		createJobResponse := &usecases.CreateJobResponse{}
		err := post("/jobs", createJobRequest, createJobResponse)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(createJobResponse.Msg)
	},
}
