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

func byteCounter(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
