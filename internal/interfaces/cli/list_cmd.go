package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var listCmd = &cobra.Command{
	Use:   "list [jobs | notifiers]",
	Short: "List jobs/notifiers",
}

var listJobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "List jobs",
	Run: func(cmd *cobra.Command, args []string) {
		findJobsResponse := uc.FindJobsResponse{}
		err := client.get("/jobs", &findJobsResponse)
		out.error(err)
		out.printf("Showing all %d jobs\n", findJobsResponse.Count)
		printFindJobsResponse(findJobsResponse)
	},
}

func printFindJobsResponse(response uc.FindJobsResponse) {
	header := []string{"NAME", "COMMAND", "TICK"}
	rows := [][]string{}
	for _, job := range response.Jobs {
		row := []string{job.Name, job.Command, job.Tick}
		rows = append(rows, row)
	}
	out.printTable(header, rows)
}

var listNotifiersCmd = &cobra.Command{
	Use:   "notifiers",
	Short: "List notifiers",
	Run: func(cmd *cobra.Command, args []string) {
		findNotifiersResponse := uc.FindNotifiersResponse{}
		err := client.get("/notifiers", &findNotifiersResponse)
		out.error(err)
		out.printf("Showing all %d notifiers\n", findNotifiersResponse.Count)
		printFindNotifiersResponse(findNotifiersResponse)
	},
}

func printFindNotifiersResponse(response uc.FindNotifiersResponse) {
	header := []string{"NAME", "TYPE"}
	rows := [][]string{}
	for _, notifier := range response.Notifiers {
		row := []string{notifier.Name, notifier.Type}
		rows = append(rows, row)
	}
	out.printTable(header, rows)
}

func init() {
	listCmd.AddCommand(listJobsCmd)
	listCmd.AddCommand(listNotifiersCmd)
}
