package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [job | channel]",
	Short: "Delete a job/channel",
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
	deleteCmd.AddCommand(deleteJobCmd)
}
