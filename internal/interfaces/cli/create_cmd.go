package cli

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create [job | channel]",
	Short: "Creates a new job/channel",
}

func init() {
	createCmd.AddCommand(createJobCmd)
}
