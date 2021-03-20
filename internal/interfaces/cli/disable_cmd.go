package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var disableCmd = &cobra.Command{
	Use:   "disable job",
	Short: "Disable a job",
}

var disableJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Disable a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var request = &uc.UpdateJobStatusRequest{Status: false}
		response := uc.UpdateJobStatusResponse{}
		err := client.put("/jobs/"+name, &request, &response)
		out.error(err)
		out.println(response.Msg)
	},
}

func init() {
	disableCmd.AddCommand(disableJobCmd)
}
