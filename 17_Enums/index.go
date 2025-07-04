package main

import "fmt"

// In go there is no enum, Need to use const to declare enum
// We can also declare custom types like myStatus
type myStatus int
type day string

// a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
// It will increment automaically
const (
	pending myStatus = iota
	progress
	fullFilled
)

const (
	sunday day = "Sunday"
	monday day = "Monday"
)

func myStatusValue(status myStatus) {
	fmt.Println("status", status)
}

func currenrDay(day day) {
	fmt.Println("day", day)
}

func main() {
	myStatusValue(pending)  // "status---->" 0
	myStatusValue(progress) // "status---->" 1

	currenrDay(sunday) // "day---->" "Sunday"
}
