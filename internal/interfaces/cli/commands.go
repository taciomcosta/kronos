package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
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
		out.printFindJobResponse(findJobsResponse)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [job | channel]",
	Short: "Delete a job/channel",
}

var deleteJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Delete a job",
	Run: func(cmd *cobra.Command, args []string) {
		deleteJobResponse := uc.DeleteJobResponse{}
		err := client.delete("/jobs/"+flags.Name, &deleteJobResponse)
		out.error(err)
		out.println(deleteJobResponse.Msg)
	},
}
