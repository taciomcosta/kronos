package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Shows execution history of all/one job",
	Run: func(cmd *cobra.Command, args []string) {
		request := parseFindExecutionsRequest(args)
		response := uc.FindExecutionsResponse{}
		url := fmt.Sprintf("/executions?jobName=%s&last=%d", request.JobName, request.Last)
		err := client.get(url, &response)
		out.error(err)
		printFindExecutionsTable(response)
	},
}

func parseFindExecutionsRequest(args []string) uc.FindExecutionsRequest {
	request := uc.FindExecutionsRequest{
		JobName: "",
		Last:    flags.Page,
	}
	if len(args) > 0 {
		request.JobName = args[0]
	}
	return request
}

func printFindExecutionsTable(response uc.FindExecutionsResponse) {
	header := []string{"NAME", "DATE", "STATUS"}
	rows := [][]string{}
	for _, exec := range response.Executions {
		row := []string{exec.JobName, exec.Date, exec.Status}
		rows = append(rows, row)
	}
	out.printTable(header, rows)
}

func init() {
	historyCmd.Flags().IntVarP(&flags.Page, "page", "p", 1, "Pagination argument")
}
