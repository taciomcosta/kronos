package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

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

var listJobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "List jobs",
	Run: func(cmd *cobra.Command, args []string) {
		findJobsResponse := uc.FindJobsResponse{}
		err := client.get("/jobs", &findJobsResponse)
		out.error(err)
		out.printf("Showing all %d jobs\n", findJobsResponse.Count)
		printFindJobResponse(findJobsResponse)
	},
}

func printFindJobResponse(response uc.FindJobsResponse) {
	header := []string{"NAME", "COMMAND", "TICK"}
	rows := [][]string{}
	for _, job := range response.Jobs {
		row := []string{job.Name, job.Command, job.Tick}
		rows = append(rows, row)
	}
	out.printTable(header, rows)
}

var deleteJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Delete a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		deleteJobResponse := uc.DeleteJobResponse{}
		err := client.delete("/jobs/"+name, &deleteJobResponse)
		out.error(err)
		out.println(deleteJobResponse.Msg)
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
}
