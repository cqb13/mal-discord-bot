package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	fmt.Printf("[%d] %s\n", time.Now().Unix(), message)
}
