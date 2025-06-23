package main

import "fmt"

func main() {
	const a string = "Test"
	fmt.Println(a)

	const b = 2
	fmt.Println(b)

	const (
		c = true
		d = 2.33
	)
	fmt.Println(c, d)
}
