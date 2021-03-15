package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var listCmd = &cobra.Command{
	Use:   "list [jobs | channels]",
	Short: "List jobs/channels",
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

func init() {
	listCmd.AddCommand(listJobsCmd)
}
