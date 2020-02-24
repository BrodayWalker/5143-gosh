package main

import (
	"strconv"
)

// Exclamation will return the n-th command
// from the commandList
func Exclamation(location string) {
	element, _ := strconv.Atoi(location)
	buffer = commandList[element]
}
