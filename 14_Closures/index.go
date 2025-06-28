package main

import "fmt"

func sum() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}

func main() {
	increment := sum()
	fmt.Println("New sum", increment())
}
