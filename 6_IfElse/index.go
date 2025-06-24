package main

import "fmt"

func main() {
	var a = 9
	if a >= 10 {
		fmt.Println("High")
	} else {
		fmt.Println("low")
	}

	// Declare variables with if statement and used with
	if b := 9; b >= 10 {
		fmt.Println("b if", b)
	} else {
		fmt.Println("b else", b)
	}
}
