package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	// Get the current working directory
	path, err := os.Getwd()

	if err != nil {
		fmt.Println("Failed to get current working directory.")
	}
	folderName := "test"

	// Construct absolute path
	totalPath := path + "\\" + folderName

	// Print path for testing
	fmt.Println(totalPath)

	// Convert string totalPath to a *uint16
	totalPathPtr := syscall.StringToUTF16Ptr(totalPath)
	// Make the directory using the Windows CreateDirectory system call
	errPath := syscall.CreateDirectory(totalPathPtr, nil)

	if errPath == syscall.ERROR_ALREADY_EXISTS {
		fmt.Println("Directory already exists.")
	} else if errPath == syscall.ERROR_PATH_NOT_FOUND {
		fmt.Println("Failed to create directory. One or more intermediate directories do not exist.")
	}
}
