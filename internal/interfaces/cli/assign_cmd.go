package cli

import (
	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var assignCmd = &cobra.Command{
	Use:   "assign [notifier] [job]",
	Short: "Assigns a notifier to a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		request := uc.AssignNotifierToJobRequest{
			NotifierName: args[0],
			JobName:      args[1],
			OnErrorOnly:  flags.OnErrorOnly,
		}
		response := uc.AssignNotifierToJobResponse{}
		err := client.post("/assignments/", &request, &response)
		out.error(err)
		out.println(response.Msg)
	},
}

func init() {
	assignCmd.Flags().BoolVarP(&flags.OnErrorOnly, "on-error-only", "", false, "Notifies only on execution errors")
}
