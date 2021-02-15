package sqlite

import (
	"errors"
	"strings"
)

var errorMap map[string]string = map[string]string{
	"1555": "Cannot create duplicate resources",
}

func parseError(err error) error {
	for code, message := range errorMap {
		if strings.Contains(err.Error(), code) {
			return errors.New(message)
		}
	}
	return err
}
