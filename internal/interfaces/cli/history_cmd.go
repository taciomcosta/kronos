package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Shows execution history of all/one job",
	Run: func(cmd *cobra.Command, args []string) {
		request := parseFindExecutionsRequest(args)
		response := uc.FindExecutionsResponse{}
		url := fmt.Sprintf("/executions?jobName=%s&page=%d", request.JobName, request.Page)
		err := client.get(url, &response)
		out.error(err)
		printFindExecutionsTable(response)
	},
}

func parseFindExecutionsRequest(args []string) uc.FindExecutionsRequest {
	request := uc.FindExecutionsRequest{
		JobName: "",
		Page:    flags.Page,
	}
	if len(args) > 0 {
		request.JobName = args[0]
	}
	return request
}

func printFindExecutionsTable(response uc.FindExecutionsResponse) {
	header := []string{"NAME", "DATE", "STATUS", "CPU TIME (ns)", "MEM USAGE (MB)"}
	rows := [][]string{}
	for _, exec := range response.Executions {
		cpu := strconv.Itoa(exec.CPUTime)
		memory := parseMemoryColumn(exec.MemUsage)
		row := []string{exec.JobName, exec.Date, exec.Status, cpu, memory}
		rows = append(rows, row)
	}
	out.printTable(header, rows)
}

func parseMemoryColumn(memoryUsage int) string {
	megabytes := float64(memoryUsage) / 1024 / 1024
	return strconv.FormatFloat(megabytes, 'f', 2, 64)
}

func init() {
	historyCmd.Flags().IntVarP(&flags.Page, "page", "p", 1, "Pagination argument")
}
