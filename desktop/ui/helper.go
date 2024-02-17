package ui

import (
	"errors"
	"fmt"
)

func toString(v any, errs ...error) (string, error) {
	return fmt.Sprintf("%v", v), errors.Join(errs...)
}

func durationCount(durationInMinutes int) string {
	if durationInMinutes < 60 {
		return fmt.Sprintf("%dm", durationInMinutes)
	}
	hours := durationInMinutes / 60
	minutes := durationInMinutes % 60
	return fmt.Sprintf("%dh%dm", hours, minutes)
}
