package cli

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete [job | channel]",
	Short: "Delete a job/channel",
}

func init() {
	deleteCmd.AddCommand(deleteJobCmd)
}
