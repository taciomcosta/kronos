package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var unassignCmd = &cobra.Command{
	Use:   "unassign [notifier] [job]",
	Short: "Unassigns a notifier from a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request := uc.UnassignNotifierFromJobRequest{
			NotifierName: args[0],
			JobName:      args[1],
		}
		response := uc.UnassignNotifierFromJobResponse{}
		err := client.delete("/assignments/", &request, &response)
		out.error(err)
		out.println(response.Msg)
	},
}

func init() {}
