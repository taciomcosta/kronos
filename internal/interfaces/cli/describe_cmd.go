package cli

import (
	"github.com/spf13/cobra"

	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var describeJobTemplate = `
Name: %s
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
	Use:   "describe [job | channel]",
	Short: "Describe a job/channel",
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

func init() {
	describeCmd.AddCommand(describeJobCmd)
}
