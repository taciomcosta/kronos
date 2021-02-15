package entities

import (
	"fmt"
	"sort"
)

// SugarExpressionMap maps sugar expressions to cron expressions
var SugarExpressionMap map[string]string = map[string]string{
	"every minute":           "* * * * *",
	"every 5 minutes":        "*/5 * * * *",
	"every 10 minutes":       "*/10 * * * *",
	"every 30 minutes":       "*/30 * * * *",
	"every hour":             "0 * * * *",
	"every noon":             "12 * * * *",
	"every midnight":         "0 0 * * *",
	"every 1st day of month": "0 1 * * *",
	"every 2 months":         "0 0 1 1-12/2 *",
	"every 3 months":         "0 0 1 1,4,7,10 *",
	"every 6 months":         "0 0 1 */6 *",
	"every year":             "0 0 1 1 *",
	"every weekday":          "0 0 * * 1-5",
	"every week":             "0 * * * 0",
	"every weekend":          "0 0 * * 6,0",
}

// GetSugarExpressions lists expressions from ExpressionMap
func GetSugarExpressions() []string {
	expressions := make([]string, 0)
	for key := range SugarExpressionMap {
		expressions = append(expressions, key)
	}
	sort.Strings(expressions)
	return expressions
}

// FormatExpression formats expression to: <cron> (<expr>)
func FormatExpression(key string) string {
	value, ok := SugarExpressionMap[key]
	if ok {
		return fmt.Sprintf("%s (%s)", value, key)
	}
	return key
}
