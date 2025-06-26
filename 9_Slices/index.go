package main

import "fmt"

func main() {
	// Slices have Dynamic size. Built on top of arrays, but can grow or shrink.
	var a = []int{1, 2, 3}
	fmt.Println("a", a)
	fmt.Println("a length", len(a))
	fmt.Println("a capacity", cap(a))
	fmt.Println("a slice", a[1:])
	fmt.Println("a slice", a[:2])

	// Shorthand syntax
	b := []int{1, 2, 3}
	fmt.Println("b", b)
	b = append(b, 4)
	fmt.Println("New b", b)

	// Use make method to create slices, make(type, size, capacity),
	// capacity: Indicates intial capacity this will gets doubled when element will be added in slices
	// Size: Initial size, Zeroed value gets added, Example: 2 ---> [0, 0]
	c := make([]int, 0, 2)
	fmt.Println("c", c)
	fmt.Println("c initial capacity", cap(c))
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)
	fmt.Println("c", c)
	fmt.Println("c new capacity", cap(c))

	// if slice is uninitialized then its nil
	var d []int
	fmt.Println("d", d)
	fmt.Println("d is nill", d == nil)
}
