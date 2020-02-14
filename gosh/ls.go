package main

import (
	"fmt"
	"io/ioutil"
)

// Ls lists all files and directories in a specified folder. Currently, no
// flags are supported.
func Ls(args []string) {
	// ReadDir reads the directory named by dirname and returns a list of
	// directory entries sorted by filename. The entries are a FileInfo
	// object with the following format:
	/*
			type FileInfo interface {
		    	Name() string       // base name of the file
		    	Size() int64        // length in bytes for regular files; system-dependent for others
		    	Mode() FileMode     // file mode bits
		    	ModTime() time.Time // modification time
		    	IsDir() bool        // abbreviation for Mode().IsDir()
				Sys() interface{}   // underlying data source (can return nil)
			}
	*/

	// If no arguments, list all file and folder names only
	if len(args) == 0 {
		files, _ := ioutil.ReadDir(".\\")
		for _, file := range files {
			fmt.Printf(file.Name() + " ")
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("No arguments to ls are currently supported.\n")
	}
}
