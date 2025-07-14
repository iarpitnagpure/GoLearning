package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file for writing (create if not exists, truncate if it does)
	// Use OpenFile instead of Open method to avoid acess denied issue
	file, err := os.OpenFile("example.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileContentSlice := []byte("Hello World")

	fileContentLength, err1 := file.Write(fileContentSlice)
	if err1 != nil {
		panic(err1)
	}

	fmt.Println("fileContentLength", fileContentLength) // "fileContentLength--->" 11
}
