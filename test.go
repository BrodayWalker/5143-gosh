package main

import (
	"bufio"
	"fmt"
	"os"
)

//list that will hold all commands typed in the terminal
var commandList []string

// will save the position of latest command saved in
// command history file
var whereLeftOff int

func main() {
	// loads the command history file contents into RAM
	loadHistory()
	commandList = append(commandList, "12345")
	commandList = append(commandList, "67890")
	saveHistory()
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
	whereLeftOff = len(commandList)
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
