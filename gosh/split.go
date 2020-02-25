package main

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

func init() {
	// Add this command's function to the command mapping
	ComMap["split"] = Split
}

var lines, size, chunks, length, counter int

//Split : split [option] [input [prefix]]
//example: split -n3 ooga.txt
func Split(args []string) {
	//loop will iterate through all elements in args
	for word := range args {
		//switch cases will only cover single-dashed flags for now
		switch string(args[word][0]) {
		case "-":
			fmt.Println(args[word])
			switch string(args[word][1]) {
			case "l":
				fmt.Println(args[word+1])
				//grabbing the number of lines of text to put in each child file
				lines, _ = strconv.Atoi(args[word+1])
				//slice off the processed flag and associated lines value
				args = args[1:]
			case "b":
				fmt.Println("This is the b flag!")
			default:
				fmt.Println("Unknown. Exiting.")
			}
		default:
			fmt.Println(args[word])
			//if the lines flag was waved
			if lines != 0 {
				fmt.Println("finna run lines code yeet")
				printByLines(lines, args[word])
			} else {
				fmt.Println("Ran default!")
				counter = 0
				file, err := os.Open(args[word]) //open file from input string array
				if err != nil {
					log.Fatal(err)
				}
				fi, _ := file.Stat()
				//this variable will hold the sections of data we will be writing to separate files
				data := make([]byte, fi.Size())
				for counter < 100 {
					//create a "sub-file" with the same name as the parent but with a counter value appended to the front
					file2, err := os.Create(strconv.Itoa(counter) + filepath.Base(args[word]))
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
