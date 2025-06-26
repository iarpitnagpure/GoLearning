package main

import (
	"fmt"
	"maps"
)

func main() {
	// In Go, a map is a built-in data type that stores key-value pairs, similar objects in JavaScript.
	// Declare using make method
	var a = make(map[string]string)
	a["name"] = "Test"
	fmt.Println("a", a) // a ---> map[name:Test]
	fmt.Println("a[name]", a["name"])
	// If key is not present then its consider as zeroed value
	fmt.Println("a[randomKey]", a["randomKey"]) // a[randomKey] ---> ""

	// Shorthand syntax b := make(map[string]int) or
	b := map[string]int{
		"name": 1,
		"age":  2,
	}
	fmt.Println("b", b)

	// Delete key
	delete(b, "name")
	fmt.Println("b after deleting name", b)

	clear(b)
	fmt.Println("b after clear", b)

	// Get element return value of key and also return boolean if key is present
	value, ok := a["name"]
	fmt.Println("value, ok", value, ok) // value ---> "Test", ok ---> true

	// Use equal method to compare two map
	m1 := map[string]int{"age": 1}
	m2 := map[string]int{"age": 1}
	fmt.Println("Is a and b is equal", maps.Equal(m1, m2)) // true
}
