package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// Can return multiple values
func newFunction() (string, int, bool) {
	return "Test", 1, true
}

func main() {
	total := sum(1, 3)
	fmt.Println("total", total)

	// access multiple values from function
	a, b, c := newFunction()
	fmt.Println("a, b, c", a, b, c)

	// functions can be assigned to variables
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("add", add(1, 3))
}
