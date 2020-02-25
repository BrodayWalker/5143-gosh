package main

import "fmt"

func init() {
	// Add this command's function to the command mapping
	ComMap["history"] = History
}

// History will print off all commands run
// since the beginning
func History(args []string) {
	for i := range commandList {
		fmt.Println(i, commandList[i])
	}
}
