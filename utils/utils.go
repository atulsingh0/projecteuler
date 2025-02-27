package utils

import (
	"fmt"
	"time"
)

func Timer(start time.Time) {
	defer func() {
		fmt.Printf("Time taken: %v\n", time.Since(start))
	}()
}
