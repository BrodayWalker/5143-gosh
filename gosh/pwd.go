package main

import (
	"fmt"
	// If not found, use command "go get golang.org/x/sys/windows"
	// sys is and alias
	sys "golang.org/x/sys/windows"
)

// Pwd prints the name of the working directory.
func Pwd(args []string) {
	if len(args) == 0 {
		// Print the working directory
		dir, err := sys.GetFullPathName(nil, nil, nil)
		fmt.Printf(dir)
	}
}
