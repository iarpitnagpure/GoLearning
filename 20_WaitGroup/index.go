package main

import (
	"fmt"
	"sync"
)

func getValues(val int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("val", val)
}

// WaitGroup is a type in Go’s sync package that is used to wait for a collection of goroutines to finish.
// Used for Launch multiple goroutines and Wait for all of them to finish before exiting the program or continuing

// Basic Workflow
// wg.Add(n) - sets the number of goroutines to wait for
// wg.Done() - Each goroutine calls defer wg.Done() when it finishes
// wg.Wait() - The main goroutine calls wg.Wait() to block until all goroutines are done

func main() {
	var wg sync.WaitGroup
	newVal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range newVal {
		wg.Add(1)
		go getValues(val, &wg) // "val--->" 9 2 1 6 4 5 7 8  3
	}
	wg.Wait()
}
