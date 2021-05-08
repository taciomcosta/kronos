package cli

import (
	"github.com/spf13/cobra"
)

var version = "0.5.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows current version",
	Run: func(cmd *cobra.Command, args []string) {
		out.println(version)
	},
}

func init() {}
