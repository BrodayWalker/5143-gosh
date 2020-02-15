package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
)

//Split : split [option] [input [prefix]]
//example: split -n3 ooga.txt
func main() {
	input := [1]string{"C:\\Users\\Owner\\Desktop\\asdfasdf\\asdf.txt"}

	//outer loop will iterate through all elements in args
	//inner loop will iterate through each character in element
	for word := range input {
		for char := range input[word] {
			//switch cases will only cover single-dashed flags for now
			switch string(input[word][char]) {
			case "-":
				flags(input[word][1:]) //process flags
			default:
				fmt.Println("Ran default!")
				counter := 42
				file, err := os.Open(input[word]) //open file from input string array
				if err != nil {
					log.Fatal(err)
				}
				//this variable will hold the sections of data we will be writing to separate files
				data := make([]byte, 100)
				for {
					//create a "sub-file" with the same name as the parent but with a counter value appended to the front
					file2, err := os.Create(strconv.Itoa(counter) + filepath.Base(input[word]))
					if err != nil {
						log.Fatal(err)
					}
					//read the data from 'file' into array 'data'
					_, err = file.Read(data)
					//write contents of 'data' to 'file2'
					file2.Write(data)
					//if we've reached the end of the file, quit looping through it
					if err == io.EOF {
						break
					}
					counter++
				}
				break
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
