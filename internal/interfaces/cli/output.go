package cli

import (
	"bytes"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/taciomcosta/kronos/internal/entities"
)

var out output

type output struct{}

func (o output) println(a ...interface{}) {
	fmt.Println(a...)
}

func (o output) printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (o output) error(err error) {
	if err != nil {
		o.println(err)
		os.Exit(1)
	}
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

func getTickExamplesTables() string {
	writer := bytes.NewBufferString("")
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"", "Examples", ""})
	table.SetAutoWrapText(false)
	table.AppendBulk(getExamples())
	table.Render()
	return writer.String()
}

func getExamples() [][]string {
	examples := [][]string{{"* * * * *", "*/2 1-12 * * *", "1,3 * * * *"}}
	expressions := entities.GetSugarExpressions()
	for i := 0; i < len(expressions); i += 3 {
		row := []string{expressions[i], expressions[i+1], expressions[i+2]}
		examples = append(examples, row)
	}
	return examples
}
