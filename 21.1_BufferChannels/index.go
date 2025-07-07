package main

import "fmt"

// Buffered channels don’t block until the buffer is full.
// To use buffer channels use stytax: myChannel := make(chan string, 2)
// Add size as second arguement in make function
func main() {
	myChannel := make(chan string, 2)
	myChannel <- "New Value"
	myChannel <- "New Value 2"
	fmt.Println("myChannel", <-myChannel) // "myChannel--->" New Value
	fmt.Println("myChannel", <-myChannel) // "myChannel--->" New Value 2
	close(myChannel)                      // close the buffer channel
}
