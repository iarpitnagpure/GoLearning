package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for index, value := range a {
		fmt.Println("index, value", index, value) // index ---> 0, value ---> 1
	}

	b := map[string]int{"phone": 1, "age": 2}
	for index, value := range b {
		fmt.Println("index, value", index, value) // index ---> "phone", value ---> 1
	}

	// If there is only one variable, It will return index
	c := map[string]int{"phone": 1}
	for index := range c {
		fmt.Println("index", index) // index ---> "phone"
	}
}
