package main

import "fmt"

// Generics allow you to write flexible and reusable code that works with different types, without sacrificing type safety.
func getValues[T int | bool](a []T) {
	for _, val := range a {
		fmt.Println("val", val)
	}
}

// Declare any type with struct
type Box[T any] struct {
	value T
}

func (b Box[T]) Get() T {
	return b.value
}

func main() {
	a := []int{1, 2, 3}
	getValues(a) // "val --->" 1 2 3

	b := []bool{true, false}
	getValues(b) // "val --->" true false

	intBox := Box[int]{value: 42}
	fmt.Println("Box", intBox.Get()) // "Box--->" 42
}
