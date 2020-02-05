package main

//For this command, the function will receive (hopefully) the portion of the command
//after 'chmod '. It will process the options by splitting the arguments string
//and then iteratively traversing each substring.

import (
	"fmt"
	"os"
	"strings"
)

//chmod u=rwx,g=rx,o=r myfile

func main() {
	initial := "577 oogabooga.exe"
	input := strings.Fields(initial) //split string apart by whitespaces
	// fmt.Println(string(input[0]))
	// fmt.Println(string(input[2]))
	// fmt.Println(string(input[7]))
	// fmt.Println(strings.Fields(input)[1])

	for _, argument := range input {
		fmt.Println("This is the argument: ", string(argument))
		for _, character := range string(argument) {
			fmt.Println(string(character))
		}
	}

	//easy part 1: if someone just typed 'chmod' with no arguments

}
