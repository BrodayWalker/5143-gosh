package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var commandList []string
	historyFile, _ := os.Open("gosh_history")
	scanner := bufio.NewScanner(historyFile)
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}
	for line := range commandList {
		fmt.Println(commandList[line])
	}
}
