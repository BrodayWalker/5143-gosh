package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var (
	//list that will hold all commands typed in the terminal
	commandList []string

	// will save the position of latest command saved in
	// command history file
	whereLeftOff int

	// ComMap with be used to create a map of strings (command keys) to Commands (the functions)
	ComMap = make(map[string]CommandFunc)
)

type (
	// CommandFunc will be used to treat functions of the given format as a type "Command"
	CommandFunc func(args []string)

	// Command : A command grouped with arguments for calling it
	Command struct {
		key  string
		args []string
	}
)

func main() {
	// loads the command history file contents into RAM
	loadHistory()
	// Main loop
	for {
		// Print prompt
		// The double percent sign must be used to print a literal percent sign
		fmt.Printf("%% ")
		// Get the user's input
		input := getInput()
		// loads a command into commandList array
		commandList = append(commandList, input)
		// Split the input by instances of && (multiple commands to run)
		commandLines := strings.Split(input, "&&")
		// Run each command line
		for _, line := range commandLines {
			// Split each line of input by pipes if necessary
			piping := strings.Split(line, "|")
			// If there is any piping, we'll need to make a Command list
			if len(piping) > 1 {
				var pipe []Command
				for _, command := range piping {
					pipe = append(pipe, parseCommand(command))
				}
				PipeLine(pipe)
			} else {
				// If there is no piping, we just need to run the command
				// Get the command from the line of text
				command := parseCommand(line)
				// Standardize command
				command.key = strings.ToLower(command.key)
				// Execute the command in the standard way
				execute(command)
			}
		}
	}
}

// loads the contents of gosh_history.tmp into the commandList array
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

func parseCommand(line string) Command {
	// Trim any leading and trailing spaces resulting from '&&' or '|' splits
	// This has to be done to process multiple commands. It just does.
	line = strings.TrimLeft(line, " ")
	line = strings.TrimRight(line, " ")
	// Separate the arguments
	symbols := strings.Split(line, " ")
	command := symbols[0]
	args := symbols[1:]
	// Return command and arguments
	return Command{command, args}
}

func execute(command Command) {
	// Route the command to call the proper function
	if com, valid := ComMap[command.key]; valid {
		com(command.args)
	} else if command.key == "exit" {
		saveHistory()
		os.Exit(0)
	} else if command.key == "test_pipe" {
		// Make a list of command lines just for testing
		comms := []Command{
			Command{
				"cat",
				[]string{"README.md"}},
			Command{
				"head",
				[]string{}},
			Command{
				"wc",
				[]string{}}}
		// Run those commands in a pipe
		PipeLine(comms)

	} else {
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
