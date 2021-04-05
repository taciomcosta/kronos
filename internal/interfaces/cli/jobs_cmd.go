package cli

import (
	"fmt"
	"strconv"
)

func parseMemory(memoryUsage int) string {
	megabytes := float64(memoryUsage) / 1024 / 1024
	return strconv.FormatFloat(megabytes, 'f', 2, 64)
}

func parseAssignedNotifiers(assignedNotifiers []string) string {
	parsed := ""
	for _, an := range assignedNotifiers {
		parsed += fmt.Sprintf("\n- %s", an)
	}
	return parsed
}
