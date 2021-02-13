package entities

import (
	"fmt"
	"strconv"
	"strings"
)

type token struct {
	values []bool
	min    int
	max    int
}

func (t *token) isSet(value int) bool {
	return t.values[value]
}

func parseToken(expression string, min int, max int) (token, error) {
	t := newToken(min, max)
	for _, s := range strings.Split(expression, ",") {
		err := t.setValues(s)
		if err != nil {
			return token{}, err
		}
	}
	return t, nil
}

func newToken(min int, max int) token {
	values := make([]bool, max+1)
	return token{values, min, max}
}

func (t *token) setValues(s string) error {
	if isStepValue(s) {
		return t.setStepValues(s)
	}
	if isRangeValue(s) {
		return t.setRangeValues(s, 1)
	}
	if isAnyValue(s) {
		t.setAnyValues(t.min, t.max, 1)
		return nil
	}
	return t.setNumericValue(s)
}

func isStepValue(value string) bool {
	return strings.Contains(value, "/")
}

func (t *token) setStepValues(stepString string) error {
	parts := strings.SplitN(stepString, "/", 2)
	step, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("can't parse %s", stepString)
	}
	if isRangeValue(parts[0]) {
		return t.setRangeValues(parts[0], step)
	}
	if isAnyValue(parts[0]) {
		t.setAnyValues(t.min, t.max, step)
		return nil
	}
	return fmt.Errorf("can't parse %s", stepString)
}

func isRangeValue(value string) bool {
	return strings.Contains(value, "-")
}

func (t *token) setRangeValues(rangeString string, step int) error {
	low, high, err := parseRange(rangeString)
	if err != nil {
		return err
	}
	if !t.isBetween(low) || !t.isBetween(high) {
		return fmt.Errorf("%s out of range %d-%d", rangeString, t.min, t.max)
	}
	t.setAnyValues(low, high, step)
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

func (t *token) isBetween(n int) bool {
	return t.min <= n && n <= t.max
}

func isAnyValue(value string) bool {
	return value == "*"
}

func (t *token) setAnyValues(min int, max int, step int) {
	for i := min; i <= max; i += step {
		t.values[i] = true
	}
}

func (t *token) setNumericValue(s string) error {
	number, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("can't parse %s", s)
	}
	if !t.isBetween(number) {
		return fmt.Errorf("%d out of range %d-%d", number, t.min, t.max)
	}
	t.values[number] = true
	return nil
}
