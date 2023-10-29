package util

import "time"

func CalculateExecutionTime(startTime time.Time, endTime time.Time) int64 {
	totalTime := endTime.Sub(startTime)
	timeInMilliseconds := totalTime.Milliseconds()
	return timeInMilliseconds
}
