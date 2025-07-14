package main

import (
	"fmt"
	"os"
)

func main() {
	// To Read file first open file using os.open method, Return file pointer and error
	file, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	// To access stats of file use .Stat(), Return file stats pointer and error
	fileStat, err1 := file.Stat()
	if err1 != nil {
		panic(err1)
	}

	// Make the buffer so that we can add file content in that variable in memory
	fileSlice := make([]byte, fileStat.Size())

	// Use .Read method on file pointer, Return the length of slice and error
	// The .Read method will add file content (In ASCII code) in buffer passed as argument
	contentLength, err2 := file.Read(fileSlice)

	if err2 != nil {
		panic(err2)
	}

	for i := 0; i < len(fileSlice); i++ {
		fmt.Println("File Content", contentLength, string(fileSlice[i]))
		// File Content 11 H
		// File Content 11 e
		// File Content 11 l
		// File Content 11 l
		// File Content 11 o
		// File Content 11
		// File Content 11 W
		// File Content 11 o
		// File Content 11 r
		// File Content 11 l
		// File Content 11 d
	}

	defer file.Close()

	// Another way to read file is using ReadFile, Return buffer slice and error
	content, err3 := os.ReadFile("example.txt")

	if err3 != nil {
		panic(err3)
	}

	fmt.Println("File Content", string(content)) // "File Content--->" Hello World

	// Note: If file is small then use Readfile otherwise use file Open approach mention in first point
}
