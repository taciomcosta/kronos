package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "kronos",
	Short: "kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs",
	Long:  "kronos is a cross-platform job scheduler that helps you manage, monitor and inspect cronjobs",
}

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(describeCmd)
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(assignCmd)
	rootCmd.AddCommand(unassignCmd)
	rootCmd.AddCommand(versionCmd)
}
