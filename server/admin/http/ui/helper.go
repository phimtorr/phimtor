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

func bytesToDisplaySize(v int64) string {
	if v < 1024 {
		return fmt.Sprintf("%d B", v)
	}
	if v < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(v)/1024)
	}
	if v < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(v)/1024/1024)
	}
	return fmt.Sprintf("%.2f GB", float64(v)/1024/1024/1024)
}
