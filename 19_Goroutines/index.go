package main

import (
	"fmt"
	"time"
)

func printValues(nums []int) {
	for _, val := range nums {
		fmt.Println("val", val)
	}
}

// Goroutines are lightweight threads managed by the Go runtime.
// They allow you to perform concurrent (non-blocking) operations easily and efficiently.
// Use "go" keyword to perform concurrent operations
func main() {
	go printValues([]int{1, 2, 3, 4, 5, 6, 7, 8}) // "val --->" 1 2 3 4 5 6 7 8

	time.Sleep(1 * time.Second) // Give goroutine time to finish
}
