package gosh

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
)

var lines, size, chunks, length, counter int

//Split : split [option] [input [prefix]]
//example: split -n3 ooga.txt
func Split(args []string) {
	input := []string{"-l", "100", "C:\\Users\\Owner\\Desktop\\asdfasdf\\asdf.txt"}
	fmt.Println(len(input))
	//loop will iterate through all elements in args
	for word := range input {
		//switch cases will only cover single-dashed flags for now
		switch string(input[word][0]) {
		case "-":
			fmt.Println(word)
			switch string(input[word][1]) {
			case "l":
				fmt.Println(input[word+1])
				//grabbing the number of lines of text to put in each child file
				lines, _ = strconv.Atoi(input[word+1])
				//slice off the processed flag and associated lines value
				input = input[1:]
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
		default:
			//if the lines flag was waved
			if lines != 0 {
				fmt.Println("finna run lines code yeet")
				printByLines(lines, input[word])
			} else {
				fmt.Println("Ran default!")
				counter = 0
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
			}
		}
	}
}

func printByLines(lineNum int, path string) {
	counter = 0                //used to name child files
	lineCounter := 0           //used to count 100 lines per child file
	file, err := os.Open(path) //open file from input string array
	if err != nil {
		log.Fatal(err)
	}
	for {
		//create a "sub-file" with the same name as the parent but with a counter value appended to the front
		file2, err := os.Create(strconv.Itoa(counter) + filepath.Base(path))
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file) //will traverse parent file
		scanner.Split(bufio.ScanLines)
		for lineCounter < 100 && scanner.Scan() { //Scan() "grabs" a line from the parent file
			file2.WriteString(scanner.Text()) //Text() returns the line from Scan() as type string
			lineCounter++
		}
		counter++
	}
}
