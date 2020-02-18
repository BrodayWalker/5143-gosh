package main

import (
	"fmt"
	"os"
	"syscall"
)

// Touch ...
func Touch(args []string) {
	// Get current working directory
	path, err := os.Getwd()
	if err == nil {
		fmt.Printf("Current path: %s\n", path)
	} else {
		fmt.Println("Error finding current working directory.")
	}

	// Make sure a filename was passed
	if len(args) == 0 {
		fmt.Println("Error: No filename name included.")
	} else {
		// Construct path + fileName
		fileName := "/" + args[0]
		path += fileName

		fmt.Printf("Path with filename: %s\n", path)

		// Hardcoded for testing only
		var perms uint32 = 0755

		// Create the file if it does not exist, do not create if the file
		// does exist
		syscall.Open(path, (os.O_CREATE | os.O_EXCL), perms)
	}
}
