package main

import "fmt"

type customer struct {
	name  string
	phone int
}

type client struct {
	id         string
	clientName string
	customer
}

// Embeded struct
func main() {
	newClient := client{
		id:         "1",
		clientName: "Test",
		customer: customer{
			name:  "Test Customer",
			phone: 9,
		},
	}
	fmt.Println("newClient", newClient) // "newClient" ---> {1 Test {Test Customer 9}}
}
