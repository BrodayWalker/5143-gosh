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
		var one, three *uint16
		var two uint32
		var four **uint16

		dir, _ := sys.GetFullPathName(one, two, three, four)
		fmt.Printf("%d", dir)
	}
}
