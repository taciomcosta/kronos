package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var describeJobTemplate = `Name: %s
Command: %s
Tick: %s
Last Execution: %s
Status: %v

Executions: 
 - %d Succeeded
 - %d Failed

Resources:
 - Average CPU: %dns
 - Average Memory: %sMB
`

var describeCmd = &cobra.Command{
	Use:   "describe [job | notifier]",
	Short: "Describe a job/notifier",
}

var describeJobCmd = &cobra.Command{
	Use:   "job",
	Short: "Describe a job",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		r := uc.DescribeJobResponse{}
		err := client.get("/jobs/"+name, &r)
		out.error(err)
		out.printf(describeJobTemplate,
			r.Name, r.Command, r.Tick, r.LastExecution, r.Status,
			r.ExecutionsSucceeded, r.ExecutionsFailed, r.AverageCPU,
			parseMemory(r.AverageMem))
	},
}

var describeNotifierTemplate = `Name: %s
Type: %s
Metadata:
%s
`

var describeNotifierCmd = &cobra.Command{
	Use:   "notifier",
	Short: "Describe a notifier",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		r := uc.DescribeNotifierResponse{}
		err := client.get("/notifiers/"+name, &r)
		out.error(err)
		out.printf(
			describeNotifierTemplate,
			r.Name, r.Type, formatMetadata(r.Metadata))
	},
}

func formatMetadata(metadata map[string]string) string {
	var formatted string
	for key, value := range metadata {
		formatted += fmt.Sprintf("- %s: %s\n", key, value)
	}
	return formatted
}

func init() {
	describeCmd.AddCommand(describeJobCmd)
	describeCmd.AddCommand(describeNotifierCmd)
}
