package main

import (
	"fmt"
	"github.com/atulsingh0/projecteuler/utils"
	"time"
)

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3,5,6,9 The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.

https://projecteuler.net/problem=1
*/

func MultipleOf3Or5(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	defer utils.Timer(time.Now())
	fmt.Println(MultipleOf3Or5(1000))
}
