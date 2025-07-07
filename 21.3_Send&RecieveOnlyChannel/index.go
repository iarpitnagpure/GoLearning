package main

import "fmt"

// Add '<-' sign before chan to make recieve only and if you add after chan keywork it will be send only
func emailSender(emailList <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for emailEntry := range emailList {
		fmt.Println("Email address is", emailEntry)
	}
}

func main() {
	emailListChannel := make(chan string, 10)
	doneChannel := make(chan bool)

	go emailSender(emailListChannel, doneChannel)

	for i := 0; i < 5; i++ {
		emailListChannel <- fmt.Sprintf("%d@gmail.com", i)
	}

	close(emailListChannel)
	<-doneChannel
}
