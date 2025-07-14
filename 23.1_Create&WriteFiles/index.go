package main

import "os"

func main() {
	// to CREATE new file use os.Create, Returns file amd error
	file, err := os.Create("example.txt")

	if err != nil {
		panic(err)
	}

	// Best practice: Always close file using defer function
	defer file.Close()

	// To write someting in file use WriteString method
	file.WriteString("New string")
}
