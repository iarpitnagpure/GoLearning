package main

import "fmt"

// num recived in function argument will be separate copy, It will allocate separate space in memory
func setNum(num int) {
	num = 5
	fmt.Println("This is value of num", num) // "num" ----> 5
}

// num recived in argument with *int will be the memory address, If we change value of that memory address it will change globally
// Append "*" to access the memory address value
func setNumWithPointer(num *int) {
	*num = 5
	fmt.Println("This is memory address of num", num) // "num" ---> 0xc00000a0d8
	fmt.Println("This is value of num", *num)         // "num" ---> 5
}

func main() {
	a := 0
	setNum(a)
	fmt.Println("value of a after function execution", a) // "a" ---> 0

	b := 0
	setNumWithPointer(&b)                                 // Need to pass memory address. Use '&' sign to get memory address
	fmt.Println("value of b after function execution", b) // "b" ---> 5
}
