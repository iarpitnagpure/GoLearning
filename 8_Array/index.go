package main

import "fmt"

func main() {
	// Array have Fixed size. The size is part of the type.
	// Declare array with length and type
	// Note: Remaning element in array will automatically assigned on zeroed value example: int--->0, bool--->false, string--->""
	// [1, 0, 0]
	var a [3]int
	a[0] = 1
	fmt.Println("a", a, len(a))

	// Declare array in one line
	var b = [3]bool{true}
	fmt.Println("b", b)

	// Declare 2D array
	var c = [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("c", c)

	// Shorthand syntax
	d := [3]int{1, 2, 3}
	fmt.Println("d", d)

	// Note: Use array when you have know array length
}
