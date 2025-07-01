package main

import (
	"fmt"
	"time"
)

type order struct {
	id     string
	amount int
	time   time.Time
}

// Use below syntax to add method on struct
// In o is first letter of order struct, Need to use * to assigned value to struct
// No need to use "*" in case of setting field, Struct automatically handle internally
func (o *order) setAmount(amount int) {
	o.amount = amount
}

// No need to use "*" operator if not setting any property
func (o order) getAmount() int {
	return o.amount
}

// Setting the constructor
func createOrder(id string, amount int, time time.Time) *order {
	testOrder := order{
		id:     id,
		amount: amount,
		time:   time,
	}

	return &testOrder
}

func main() {
	// If any property is unassigned ut will automatically assigned as zeroed value
	newOrder := order{
		id:   "1",
		time: time.Now(),
	}
	fmt.Println("newOrder", newOrder) // "newOrder" ----> {1 0 {13984000570978060760 522601 0x4b3980}}

	newOrder.setAmount(10)
	fmt.Println("newOrder after setAmount", newOrder)                             // "newOrder" ----> {1 10 {13984001168311452804 1 0x203980}}
	fmt.Println("newOrder access only id direcly", newOrder.id)                   // "id" ---> 1
	fmt.Println("newOrder access only amount using method", newOrder.getAmount()) // "Amount" ---> 10

	// Shorthand syntax
	car := struct {
		model  string
		amount int
	}{
		model:  "Nike",
		amount: 100,
	}
	fmt.Println("car", car) // "car" ----> {Nike 100}

	// constructor use
	newOrder2 := createOrder("1", 10, time.Now())
	fmt.Println("newOrder2", newOrder2) // "newOrder2" ----> &{1 10 {13984002094498374116 1369401 0xfa4980}}
}
