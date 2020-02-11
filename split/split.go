package main

import (
	"fmt"
	// "log"
	// "strings"
)

//Split : split [option] [input [prefix]]
//example: split -n3 ooga.txt
func main() {
	input := [1]string{"asdf.txt"}

	//outer loop will iterate through all elements in args
	//inner loop will iterate through each character in element
	for word := range input {
		fmt.Println(input[word] + " ")
		for char := range input[word] {
			fmt.Println(string(input[word][char]) + " ")

			//switch cases will only cover single-dashed flags for now
			switch string(input[word][char]) {
			case "-":
				flags(input[word][1:])
			}
		}
	}
}

func flags(arg string) {
	switch arg {
	case "l":
		fmt.Println("This is the l flag!")
	case "b":
		fmt.Println("This is the b flag!")
	case "C":
		fmt.Println("This is the C flag!")
	case "n":
		fmt.Println("This is the n flag!")
	case "a":
		fmt.Println("This is the a flag!")
	case "d":
		fmt.Println("This is the d flag!")
	case "x":
		fmt.Println("This is the x flag!")
	case "t":
		fmt.Println("This is the t flag!")
	case "u":
		fmt.Println("This is the u flag!")
	default:
		fmt.Println("Unknown. Exiting.")
	}
}
