package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var createCmd = &cobra.Command{
	Use:   "create [job | notifier]",
	Short: "Creates a new job/notifier",
}

var createJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Creates a new job",
	Run: func(cmd *cobra.Command, args []string) {
		createJobRequest := uc.CreateJobRequest{
			Name:    flags.Name,
			Command: flags.Command,
			Tick:    flags.Tick,
		}
		createJobResponse := &uc.CreateJobResponse{}
		err := client.post("/jobs", createJobRequest, createJobResponse)
		out.error(err)
		out.println(createJobResponse.Msg)
	},
}

func init() {
	createJobCmd.Flags().StringVarP(&flags.Name, "name", "n", "", "Unique job name")
	createJobCmd.Flags().StringVarP(&flags.Command, "cmd", "c", "", "Job entrypoint")
	createJobCmd.Flags().StringVarP(
		&flags.Tick,
		"tick", "t", "",
		"Cron expression or sugar expression:\n"+getTickExamplesTables(),
	)
	_ = createJobCmd.MarkFlagRequired("name")
	_ = createJobCmd.MarkFlagRequired("cmd")
	_ = createJobCmd.MarkFlagRequired("tick")

	createCmd.AddCommand(createJobCmd)
}
