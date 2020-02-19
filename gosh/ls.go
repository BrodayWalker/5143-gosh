package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func init(){
    // Add this command's function to the command mapping
    ComMap["ls"] = Ls
}

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
		files, err := ioutil.ReadDir(".")
        if err != nil{
            fmt.Println(err)
        }
		for _, file := range files {
			fmt.Printf(file.Name() + " ")
		}
		fmt.Printf("\n")
	} else {
		// Handle flags
		// Compile the expression to test against
		// A flag must be in args[0] or it will not be caught
		var flags, _ = regexp.Compile(`^-{1,2}(\d|\D)`)
		if flags.MatchString(args[0]) {
			// Print flags for testing
			fmt.Printf("Flags: ")
			for i := 0; i < len(args); i++ {
				fmt.Printf(args[i])
			}
			fmt.Printf("\n")
			// -l
			var l, _ = regexp.Match("l", []byte(args[0]))
			if l {
				files, _ := ioutil.ReadDir(".")
				for _, file := range files {
					// Print permissions
					fmt.Printf("%s ", file.Mode())
					// Print owner fields

					// Print size
					fmt.Printf("%d ", file.Size())
					// Print date
					fmt.Printf("%v ", file.ModTime())
					// Print file/folder name
					fmt.Printf(file.Name() + " ")
					fmt.Printf("\n")
				}
			}
		} else {

		}
	}
}
