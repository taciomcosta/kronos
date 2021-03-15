package cli

import "strconv"

func parseMemory(memoryUsage int) string {
	megabytes := float64(memoryUsage) / 1024 / 1024
	return strconv.FormatFloat(megabytes, 'f', 2, 64)
}
