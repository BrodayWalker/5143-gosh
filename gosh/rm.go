package main

import (
	"fmt"
	"os"
)

// Rm removes a folder or file
func Rm(args []string) {
	if len(args) != 0 {
		rmMe := args[0]
		err := os.Remove(rmMe)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
