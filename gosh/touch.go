package main

import (
	"fmt"
	//"os"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["touch"] = Touch
}

// Touch ...
func Touch(args []string) {
	path := args[0]
	exists := FileExists(path)
	fmt.Printf("Answer: %v\n", exists)

	/* TESTING

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

		// Open
		fp, createErr := os.Create(path)
		if createErr == nil {
			fp.Close()
		} else {
			fmt.Println("Error creating file.")
		}
	}

	END TESTING
	*/
}
