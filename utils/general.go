package utils

import (
	"fmt"
	"strconv"
	"time"
)

var StartTime time.Time

func CalculateUptime() string {
	timeSinceStart := time.Since(StartTime)

	totalSeconds := int(timeSinceStart.Seconds())

	weeks := totalSeconds / (7 * 24 * 3600)
	totalSeconds %= 7 * 24 * 3600

	days := totalSeconds / (24 * 3600)
	totalSeconds %= 24 * 3600

	hours := totalSeconds / 3600
	totalSeconds %= 3600

	minutes := totalSeconds / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%dw %dd %dh %dm %ds", weeks, days, hours, minutes, seconds)
}

func Ternary(cond bool, a string, b string) string {
	if cond {
		return a
	}
	return b
}

func UnixStampStrToPrettyStr(str string) (string, error) {
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return "", err
	}

	time := time.Unix(number, 0)

	return fmt.Sprintf("%d/%d/%d | %d:%d:%d", time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute(), time.Second()), nil
}

func RFC3339StrToPrettyStr(str string) (string, error) {
	time, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d/%d/%d | %d:%d:%d", time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute(), time.Second()), nil
}

func TimeToPrettyStr(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d | %d:%d:%d", t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute(), t.Second())
}
