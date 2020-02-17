package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	// Main loop
	for {
		// Print prompt
		// The double percent sign must be used to print a literal percent sign
		fmt.Printf("%% ")
		// Get the command
		command, args := parseCommand(getInput())
		// Standardize command
		command = strings.ToLower(command)
		// Match a function
		switchboard(command, args)
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

	// Trim \r\n for Windows
	if runtime.GOOS == "windows" {
		line = strings.TrimRight(line, "\r\n")
	} else {
		line = strings.TrimRight(line, "\n")
	}

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

func switchboard(command string, args []string) {
	// Route the command to call the proper function
	// All commands are in alphabetical order
	switch command {
	case "echo":
		Echo(args)
	case "exit":
		os.Exit(0)
	case "head":
		Head(args)
	case "ls":
		Ls(args)
	case "mkdir":
		Mkdir(args)
	case "mv":
		Mv(args)
	case "pwd":
		Pwd(args)
	case "rm":
		Rm(args)
	case "touch":
		Touch(args)
	case "wc":
		Wc(args)
	default:
		fmt.Println("Command not found.")
	}
}
