package main

import (
	"fmt"
	"github.com/atulsingh0/projecteuler/euler"
	"time"
)

func timer(start time.Time) {
	defer func() {
		fmt.Printf("Time taken: %v\n", time.Since(start))
	}()
}

func main() {

	defer timer(time.Now())
	fmt.Println(euler.MultipleOf3Or5(1000))
}
