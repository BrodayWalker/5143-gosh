package gosh

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Wc(args []string) {

	// Files object to store content
	files := make(map[string]string)
	for _, arg := range args {
		// If the argument is not a flag
		if arg[0] != '-' {
			// Read each file
			content, err := ioutil.ReadFile(arg)
			// Check for errors
			if err != nil {
				log.Fatal(err)
			}
			// Store the files contents
			files[arg] = string(content)
		}
	}

	// Let's start printing them
	for key, val := range files {
		lines := strconv.Itoa(len(strings.SplitN(val, "\n", -1)))
		words := strconv.Itoa(len(strings.SplitN(val, " ", -1)))
		chars := strconv.Itoa(len(strings.SplitN(val, "", -1)))
		fmt.Print(lines + " ")
		fmt.Print(words + " ")
		fmt.Print(chars + " ")
		fmt.Println(key)
	}

}
