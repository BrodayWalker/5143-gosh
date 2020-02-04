package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Main loop
	for {
		// Get the command
		command, args := parseCommand(getInput())

		// Temporary testing. Prints command & args
		fmt.Println(command)
		for _, element := range args {
			fmt.Println(element)
		}
	}
}

func getInput() string {
	// Create a keyboard reader
	keyboard := bufio.NewReader(os.Stdin)
	// Read a line of input
	line, e := keyboard.ReadString('\n')
	// Print out any errors
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	line = strings.Trim(line, "\n")
	return line
}

func parseCommand(line string) (string, []string) {
	// Separate the arguments
	input := strings.Split(line, " ")
	command := input[0]
	args := input[1:]
	// Return command and arguments
	return command, args
}
