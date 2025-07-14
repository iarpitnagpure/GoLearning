package main

import "os"

func main() {
	// To remove file use .Remove method, Returns error
	os.Remove("example.txt")
}
