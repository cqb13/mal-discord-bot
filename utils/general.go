package utils

import (
	"fmt"
	"strconv"
	"time"
)

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
