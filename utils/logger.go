package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	time := time.Now()
	fmt.Printf("[%02d/%02d/%d | %02d:%02d:%02d] %s\n", time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute(), time.Second(), message)
}
