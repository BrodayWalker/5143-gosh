package main

import (
	"fmt"
	sys "golang.org/x/sys/windows"
)

// Mv ...
func Mv(args []string) {
	// There is a current file (args[0]) and a target name (args[1])
	if len(args) == 2 {
		_ = sys.Rename(args[0], args[1])
	} else if len(args) == 1 {
		fmt.Println("Not enough arguments. No mv done.")
	} else {
		fmt.Println("Too many arguments.")
	}
}
