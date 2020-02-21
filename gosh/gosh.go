package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

//list that will hold all commands typed in the terminal
var commandList []string

func main() {
	// load the commands history into RAM
	loadHistory()
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

func loadHistory() {
	historyFile, _ := os.Open("gosh_history")
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}
	for line := range commandList {
		fmt.Println(commandList[line])
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

	//append command to the commandList
	commandList = append(commandList, line)

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
	case "cat":
		Echo(args)
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
	case "split":
		Split(args)
	case "touch":
		Touch(args)
	case "wc":
		Wc(args)
	default:
		fmt.Println("Command not found.")
	}
}
