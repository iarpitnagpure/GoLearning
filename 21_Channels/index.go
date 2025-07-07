package main

import "fmt"

func getChannelValue(myChannel chan int) {
	// defer function is clean up function, This will call once evrything is executed in current function
	defer func() {
		myChannel <- 10
	}()

	fmt.Println("Processed getChannelValue function")
}

// In Go, a channel is a typed conduit through which goroutines communicate.
// They are essential for synchronization and data sharing in concurrent Go programs.
// Syntax: make(chan Type) This creates a bidirectional channel of the given type.
// newChannel <- "Test" this will send values to channel
// myChannelValue := <-newChannel this will recieve value from channel
func main() {
	// Below example is typical case of deadlock if channel used without goroutines/main thread
	// newChannel := make(chan string)
	// newChannel <- "Test" // fatal error: all goroutines are asleep - deadlock!
	// myChannelValue := <-newChannel
	// fmt.Println("myChannelValue", myChannelValue)

	newChannel := make(chan int)
	go getChannelValue(newChannel)          // This will block the thread till the values are send in channel
	fmt.Println("newChannel", <-newChannel) // "Processed getChannelValue function", "newChannel--->" 10
	// Channel is also used instaed of waitgroup to block the operation till all goroutines are done.
	// Tip: Use Channel as blocking when you have single goroutines, If you have multiple goroutines then use waitgroup to block the thread
}
