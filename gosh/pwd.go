package main

import (
	"fmt"
	"os"
)

// Pwd prints the name of the working directory.
func Pwd(args []string) {
	if len(args) == 0 {
		dir, err := os.Getwd()
		if err == nil {
			fmt.Println(dir)
		}
	}
}
