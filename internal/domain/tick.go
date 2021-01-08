package domain

import (
	"fmt"
	"strconv"
	"strings"
)

func parseToken(token string, min int, max int) ([]int, error) {
	values := make([]bool, max+1)
	for _, s := range strings.Split(token, ",") {
		err := setValues(values, s, min, max)
		if err != nil {
			return []int{}, err
		}
	}
	return filterSetValues(values), nil
}

func setValues(values []bool, s string, min int, max int) error {
	if isStepValue(s) {
		return setStepValues(values, s, min, max)
	}
	if isRangeValue(s) {
		return setRangeValues(values, s, min, max, 1)
	}
	if isAnyValue(s) {
		setAnyValues(values, min, max, 1)
		return nil
	}
	return setNumericValue(values, s, min, max)
}

func isStepValue(value string) bool {
	return strings.Contains(value, "/")
}

func setStepValues(values []bool, stepString string, min int, max int) error {
	parts := strings.SplitN(stepString, "/", 2)
	step, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("can't parse %s", stepString)
	}
	if isRangeValue(parts[0]) {
		return setRangeValues(values, parts[0], min, max, step)
	}
	if isAnyValue(parts[0]) {
		setAnyValues(values, min, max, step)
		return nil
	}
	return fmt.Errorf("can't parse %s", stepString)
}

func isRangeValue(value string) bool {
	return strings.Contains(value, "-")
}

func setRangeValues(values []bool, rangeString string, min, max, step int) error {
	low, high, err := parseRange(rangeString)
	if err != nil {
		return err
	}
	if !isBetween(low, min, max) || !isBetween(high, min, max) {
		return fmt.Errorf("%s out of range %d-%d", rangeString, min, max)
	}
	setAnyValues(values, low, high, step)
	return nil
}

func parseRange(rangeString string) (int, int, error) {
	rangeValues := strings.SplitN(rangeString, "-", 2)
	low, errLow := strconv.Atoi(rangeValues[0])
	high, errHigh := strconv.Atoi(rangeValues[1])
	if errLow != nil || errHigh != nil {
		return 0, 0, fmt.Errorf("missing value in range %s", rangeString)
	}
	return low, high, nil
}

func isBetween(n int, min int, max int) bool {
	return min <= n && n <= max
}

func isAnyValue(value string) bool {
	return value == "*"
}

func setAnyValues(values []bool, min int, max int, step int) {
	for i := min; i <= max; i += step {
		values[i] = true
	}
}

func setNumericValue(values []bool, s string, min int, max int) error {
	number, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("can't parse %s", s)
	}
	if !isBetween(number, min, max) {
		return fmt.Errorf("%d out of range %d-%d", number, min, max)
	}
	values[number] = true
	return nil
}

func filterSetValues(values []bool) []int {
	var result []int
	for i, isSet := range values {
		if isSet {
			result = append(result, i)
		}
	}
	return result
}
