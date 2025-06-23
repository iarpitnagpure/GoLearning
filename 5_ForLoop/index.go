package main

import "fmt"

func main() {
	// while Loop
	i := 1
	for i <= 3 {
		fmt.Println("For used as While Loop", i)
		i++
	}

	// For Loop
	for i := 0; i <= 3; i++ {
		fmt.Println("For Loop", i)
	}

	// continue will not excecute that iteration, 2 will be skipped
	for i := 0; i <= 3; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Continue statement", i)
	}

	// break will exit from that loop
	for i := 0; i <= 3; i++ {
		if i == 2 {
			break
		}
		fmt.Println("Break statement", i)
	}

	// Range will iterate till n-1, start with 0
	for i := range 3 {
		fmt.Println("Range", i)
	}
}
