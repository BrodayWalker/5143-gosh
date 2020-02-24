package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var
//list that will hold all commands typed in the terminal
var commandList []string

// will save the position of latest command saved in
// command history file
var whereLeftOff int

func main() {
	// loads the command history file contents into RAM
	loadHistory()
	// Main loop
	for {
		// Print prompt
		// The double percent sign must be used to print a literal percent sign
		fmt.Printf("%% ")
		// Get the command
		line := getInput()
		// loads a command into commandList array
		commandList = append(commandList, line)
		//parse the command by dividing the arguments and the command word
		command, args := parseCommand(line)
		// Standardize command
		command = strings.ToLower(command)
		// Match a function
		switchboard(command, args)
	}
}

// loads the contents of gosh_history.tmp into the commandList array
//
func loadHistory() {
	historyFile, _ := os.Open("gosh_history.tmp")
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}
	historyFile.Close()
	// remember where the commands for this current session begin
	whereLeftOff = len(commandList)
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
		//upload command history from commandList to gosh_history.tmp
		saveHistory()
		os.Exit(0)
	case "head":
		Head(args)
	case "history":
		History()
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

func saveHistory() {
	limit := len(commandList)
	historyFile, _ := os.OpenFile("gosh_history.tmp", os.O_WRONLY|os.O_APPEND, 0)
	// add only the commands executed during this session to the
	// gosh history file by starting from the point we saved in
	// whereLeftOff
	for i := whereLeftOff; i < limit; i++ {
		historyFile.WriteString("\n" + commandList[i])
	}
	historyFile.Close()
}
