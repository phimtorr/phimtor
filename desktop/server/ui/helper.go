package ui

import (
	b64 "encoding/base64"
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

func toBase64(data []byte) string {
	return b64.StdEncoding.EncodeToString(data)
}

func toBase64Src(mineType string, content []byte) string {
	return fmt.Sprintf("data:%s;base64,%s", mineType, toBase64(content))
}
