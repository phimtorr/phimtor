package ui

import (
	"fmt"
	"time"
)

func toString(i any) string {
	return fmt.Sprintf("%v", i)
}

func secondsToDisplayTime(v any) string {
	m, ok := v.(float64)
	if !ok {
		return ""
	}
	return time.Unix(int64(m), 0).Format("2006-01-02 15:04:05")
}

func millisecondsToDisplayTime(v int64) string {
	return time.Unix(v/1000, 0).Format("2006-01-02 15:04:05")
}
