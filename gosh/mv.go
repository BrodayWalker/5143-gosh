// +build windows

package gosh

import (
	"fmt"
	sys "golang.org/x/sys/windows"
	"os"
)

// Mv ...
func Mv(args []string) {
	// There is a current file (args[0]) and a target name (args[1])
	if len(args) == 2 {

		// This just renames some file (if it exists) in the current directory
		// to a new name
		wd, pathErr := os.Getwd()
		if pathErr != nil {
			fmt.Printf("%v\n", pathErr)
		} else {
			oldPath := wd + "\\" + args[0]
			newPath := wd + "\\" + args[1]
			fmt.Println(oldPath)
			fmt.Println(newPath)
			renameErr := sys.Rename(oldPath, newPath)
			if renameErr != nil {
				fmt.Printf("%v\n", renameErr)
			}
		}
	}
}
