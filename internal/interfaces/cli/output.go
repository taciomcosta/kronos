package cli

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	uc "github.com/taciomcosta/kronos/internal/usecases"
)

var out output

type output struct{}

func (o output) println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

func (o output) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func (o output) error(err error) {
	if err != nil {
		o.println(err)
		os.Exit(1)
	}
}

func (o output) printFindJobResponse(response uc.FindJobsResponse) {
	header := []string{"NAME", "COMMAND", "TICK"}
	rows := [][]string{}
	for _, job := range response.Jobs {
		row := []string{job.Name, job.Command, job.Tick}
		rows = append(rows, row)
	}
	o.printTable(header, rows)
}

func (o output) printTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(rows)
	table.Render()
}
