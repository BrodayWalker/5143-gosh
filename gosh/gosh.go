package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Treat functions of the given format as a type "Command"
type CommandFunc func(args []string)

// Create a map of strings (command keys) to Commands (the functions)
var ComMap = make(map[string]CommandFunc)

// A command grouped with arguments for calling it
type Command struct {
    key string
    args []string
}

func main() {
	// Main loop
	for {
		// Print prompt
		// The double percent sign must be used to print a literal percent sign
		fmt.Printf("%% ")
		// Get the command
		command := parseCommand(getInput())
		// Standardize command
		command.key = strings.ToLower(command.key)
		// Match a function
		execute(command)
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

func parseCommand(line string) Command {
	// Separate the arguments
	input := strings.Split(line, " ")
	command := input[0]
	args := input[1:]
	// Return command and arguments
	return Command{command, args}
}

func execute(command Command) {
    // Route the command to call the proper function
    if com, valid := ComMap[command.key]; valid{
        com(command.args)
    }else if command.key == "exit"{
        os.Exit(0)
    }else if command.key == "test_pipe"{

        // Make a list of command lines just for testing
        comms := []Command{
            Command{
                "cat",
                []string{"README.md"} },
            Command{
                "head",
                []string{} },
            Command{
                "wc",
                []string{} } }
        // Run those commands in a pipe
        PipeLine(comms)

    }else{
        fmt.Println("Command not found.")
    }
}