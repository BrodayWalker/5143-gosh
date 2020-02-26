package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// will be used to count number of instances of the first character of each line
var numDict = map[string]int{"0": 0, "1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0}
var lowerDict = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
var upperDict = map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "E": 0, "F": 0, "G": 0, "H": 0, "I": 0, "J": 0, "K": 0, "L": 0, "M": 0, "N": 0, "O": 0, "P": 0, "Q": 0, "R": 0, "S": 0, "T": 0, "U": 0, "V": 0, "W": 0, "X": 0, "Y": 0, "Z": 0}

// // a cell will hold a line's first character and its line number
// type cell struct {
// 	char    string
// 	lineNum int
// }

// // will be sorted in place of sorting the file itself
// var fileArray []cell

func init() {
	// Add this command's function to the command mapping
	ComMap["sort"] = Sort
}

// Sort ..
// 1. I want to implement a dictionary {key:value} where the key is a character found in the input file to be sorted, and value
// 		is the number of occurances of that key in the file. Keys are only the first characters of each line.
// 2. sortFile will first run through the entire input file and update the dictionary.
// 3. Create an array of structs, where the struct is a string and an int: the first character of a line and the line number, respectively
// 4. Sort the array of structs by the character
// 5. Now, for each series of elements where the characters are the same, replace the struct's characters with the next character in the line
// 6. Sort the sub array, repeating 5 and 6 until that entire subarray is sorted.
// 7. Move on to the next sub array that have the same characters, repeating 5 through 7 until the entire array is sorted.
// 8. Print, using the line numbers of the struct, the sorted lines to Stdout.
func Sort(args []string) {
	argsSize := len(args)
	if argsSize == 1 {
		sortFile(args[0]) // call sorting function with input file; standard procedure
	} else if argsSize > 1 {
		args = append(args, ">") // append the redirection arrow and a temporary file to args
		args = append(args, "sort.tmp")
		Cat(args)                  // have Cat concatenate all the files and stick them into the sort.tmp file
		sortFile(args[argsSize-1]) // call sorting function with temporary file, sort.tmp
	} else {
		fmt.Println("Error: no input arguments. Use form 'sort input1 input2 inputN'")
	}
}

func sortFile(file string) {
	var (
		numbers, lowers, uppers, symbols = []string{}, []string{}, []string{}, []string{}
		element                          int
	)
	inputFile, _ := os.OpenFile(file, os.O_RDONLY, 0755)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if scanner.Text() != "" {
			if _, numIn := numDict[string(scanner.Text()[0])]; numIn {
				numbers = append(numbers, scanner.Text())
			} else if _, lowerIn := lowerDict[string(scanner.Text()[0])]; lowerIn {
				lowers = append(lowers, scanner.Text())
			} else if _, upperIn := upperDict[string(scanner.Text()[0])]; upperIn {
				uppers = append(uppers, scanner.Text())
			} else {
				symbols = append(symbols, scanner.Text())
			}
		}
	}
	os.Remove("sort.tmp")
	inputFile.Close()
	sort.Strings(symbols)
	sort.Strings(numbers)
	sort.Strings(lowers)
	sort.Strings(uppers)
	for element = range symbols {
		fmt.Println(symbols[element])
	}
	for element = range numbers {
		fmt.Println(numbers[element])
	}
	for element = range lowers {
		fmt.Println(lowers[element])
	}
	for element = range uppers {
		fmt.Println(uppers[element])
	}
}

// for ele := range args {
// 	fmt.Println(args[ele])
// }
