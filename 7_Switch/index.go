package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 10
	switch a {
	case 10:
		fmt.Println("Case", a)
	default:
		fmt.Println("Default", a)
	}

	// Multiple Condition Switch
	// Add multiple conditions on case with ,
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its Weekend")
	default:
		fmt.Println("Its Weekday")
	}

	// Type Switch
	// i is any when declare with interface{}, i.(type) is
	newFunc := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("This is int", i)
		case bool:
			fmt.Println("This is boolean", i)
		}
	}
	newFunc(1)
	newFunc(false)
}
