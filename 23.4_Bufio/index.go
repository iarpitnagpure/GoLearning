package main

import (
	"bufio"
	"os"
)

// Using bufio we can read write data on file on stream fashion
func main() {
	// Open new soruce file
	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	// Create new source destination file
	destinationFile, err1 := os.Create("example1.txt")

	if err1 != nil {
		panic(err1)
	}

	defer destinationFile.Close()

	// Create the reader and writer to read and write data from source to destination file
	sourceFileBuf := bufio.NewReader(sourceFile)
	destinationFileBuf := bufio.NewWriter(destinationFile)

	for {
		// Read data using reader, Returns bytes and error
		newByte, err2 := sourceFileBuf.ReadByte()

		if err2 != nil {
			if err2.Error() != "EOF" {
				panic(err2)
			}

			break
		}

		// Write data using writer, Returns error
		destinationFileBuf.WriteByte(newByte)
	}

	destinationFileBuf.Flush()
}
