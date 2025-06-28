package main

import "fmt"

// A variadic function is a function that accepts zero or more arguments of a specific type.
// It uses ... before the type in the parameter list.
func sum(num ...int) int {
	total := 0
	for _, value := range num {
		total = total + value
	}

	return total
}

func main() {
	a := []int{1, 3, 4}
	// use a... to spread number as parameters
	result := sum(a...)
	fmt.Println("result", result)
}
