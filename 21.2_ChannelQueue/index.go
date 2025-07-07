package main

import "fmt"

// Recived two channel emailList will recieve email list in another goroutine
// done channel value gets set when function opetation is complete
func emailSender(emailList chan string, done chan bool) {
	defer func() { done <- true }()

	for emailEntry := range emailList {
		fmt.Println("Email address is", emailEntry)
	}
}

// emailListChannel is buffer channel with size 10, If email is above 10 then it will block the other operations
// doneChannel is normal channel will pass to emailSender function and set the value in clean up function
// emailSender function will add in separate goroutines and listen to emailListChannel channel
func main() {
	emailListChannel := make(chan string, 10)
	doneChannel := make(chan bool)

	go emailSender(emailListChannel, doneChannel)

	for i := 0; i < 5; i++ {
		emailListChannel <- fmt.Sprintf("%d@gmail.com", i)
	}

	// close the emailListChannel buffeer channel to avaoid any deadloock error
	close(emailListChannel)
	// Once doneChannel will recieved value the program will close
	<-doneChannel
}
