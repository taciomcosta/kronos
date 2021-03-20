package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var enableCmd = &cobra.Command{
	Use:   "enable job",
	Short: "Enable a job",
}

var enableJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Enable a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		var request = &uc.UpdateJobStatusRequest{Status: true}
		response := uc.UpdateJobStatusResponse{}
		err := client.put("/jobs/"+name, &request, &response)
		out.error(err)
		out.println(response.Msg)
	},
}

func init() {
	enableCmd.AddCommand(enableJobCmd)
}
