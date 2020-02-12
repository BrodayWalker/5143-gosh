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
		// Replace these
		var a *uint16
		// Print the working directory
<<<<<<< HEAD
		var one, three *uint16
		var two uint32
		var four **uint16

		dir, _ := sys.GetFullPathName(one, two, three, four)
		fmt.Printf("%d", dir)
=======
		dir, err := sys.GetFullPathName(a, 4, nil, nil)
		fmt.Println(dir)
		fmt.Println(err)
>>>>>>> cf11578b0429f4ce410d719b9148bc1c4fffec14
	}
}
