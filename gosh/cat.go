package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

//these flags will be set off if a > or >> are included in the Cat command
var (
	single, double bool = false, false

	// scanner pointer that will point to a file to be read from
	scanner *bufio.Scanner

	// this string will save the path of the file to which output will be redirected when > or >> are used
	outputPath string
)

func init() {
	// Add this command's function to the command mapping
	ComMap["cat"] = Cat
}

// Cat will either concatenate a file and print it to std out or
// intake two or more files and print them as if they were concatenated.
// One could also redirect the output to a file.
// Usage:
// 			Cat file
// 			Cat file1 file2 fileN
//func Cat(args []string)
func Cat(args []string) {
	//base case, where there are no arguments to the command
	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for {
			ctext, _ := reader.ReadString('\n')
			//exit condition faulty. Needs more testing at a later date. Use CTRL+C for now.
			if ctext == "q" {
				break
			}
			fmt.Println(ctext)
		}
		//this will run if the Cat command was run with only one argument - a file to print to std out
	} else {
		tempFile, err := os.Create("temp.txt") //create a temporary file to store input file data in
		if err != nil {
			log.Fatal(err)
		}
	mainLoop: //this is a label which I use to specify what loop, switch, or select I want to affect
		for element := range args {
			switch args[element] {
			case ">":
				//redirect to a file, erasing file's existing data if any existed
				single = true
				outputPath = args[element+1] //remember the path of the redirected output file
				break mainLoop               //kill the loop, not the switch (which is what it would do if the instruction was just "break")
			case ">>":
				//append output to a file
				double = true
				outputPath = args[element+1]
				break mainLoop
			default:
				file, _ := os.Open(args[element]) //open the file
				scanner = bufio.NewScanner(file)  //create a scanner to run through the file
				scanner.Split(bufio.ScanLines)    //ascribe the delimiter to lines, meaning the scanner will consider one line is a token
				for scanner.Scan() {              //run through the file and print each line to the stdout
					tempFile.WriteString(scanner.Text() + "\n") //write, instead of straight to stdout or a file, to a temp file
				}
				file.Close()
			}
		}
		tempFile.Seek(0, io.SeekStart) //return scanner to the top of tempFile
		scanner = bufio.NewScanner(tempFile)
		scanner.Split(bufio.ScanLines)
		if single == true {
			//opens the read-write file and truncates if successful. Otherwise, it creates the file with rw-rw-rw permissions
			outfile, _ := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			for scanner.Scan() {
				outfile.WriteString(scanner.Text() + "\n") // output contents of tempFile to outfile
			}
			outfile.Close()
		} else if double == true {
			//opens the read-write file and truncates if successful. Otherwise, it creates the file with rw-rw-rw permissions
			outfile, _ := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE, 0666)
			outfile.Seek(0, io.SeekEnd) //point at the end of outfile so we can start printing from tempFile at that point in outfile (appending!)
			for scanner.Scan() {
				outfile.WriteString("\n" + scanner.Text()) // output contents of tempFile to outfile
			}
			outfile.Close()
		} else {
			// print the contents of tempFile to stdout
			for scanner.Scan() {
				fmt.Println(scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		//close and delete the temporary file
		tempFile.Close()
		err = os.Remove("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
	}
}
