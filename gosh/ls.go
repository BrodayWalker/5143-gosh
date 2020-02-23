package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func init() {
	// Add this command's function to the command mapping
	ComMap["ls"] = Ls
}

// Ls lists all files and directories in a specified folder. Currently, no
// flags are supported.
func Ls(args []string) {
	argList, flags := ArgSplitter(args)
	var path string = ""

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
	if len(argList) == 0 && len(flags) == 0 {
		defaultPrint(".")
	} else if len(argList) > 0 && len(flags) == 0 {
		// Different path than current working directory, no flags passed
		path = buildPath(argList[0])
		defaultPrint(path)
	} else if len(argList) == 0 && len(flags) > 0 {
		longPrint(".")
	} else if len(argList) > 0 && len(flags) > 0 {
		path = buildPath(argList[0])
		for _, v := range flags {
			if v == "l" {
				longPrint(path)
			}
		}
	}
}

func buildPath(path string) string {
	absPath, fpErr := filepath.Abs(path)
	// Print error
	if fpErr != nil {
		fmt.Println("Absolute path error")
	}
	// Get FileInfo struct for our path
	pathInfo, statErr := os.Stat(absPath)
	// If path is not valid
	// Path must exist and must be a directory
	if statErr != nil || pathInfo.IsDir() == false {
		fmt.Println(statErr)
		return ""
	}

	return absPath
}

func defaultPrint(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Printf(file.Name() + " ")
	}
	fmt.Printf("\n")
}

func longPrint(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		// Print permissions
		fmt.Printf("%s ", file.Mode())
		// Print owner fields
		// This is not implemented as Windows returns -1 for
		// the group and owner fields
		// Print size
		fmt.Printf("%12d ", file.Size())
		// Print date
		t := file.ModTime()
		fmt.Printf("%v ", t.Format(time.UnixDate))
		// Print file/folder name
		fmt.Printf(file.Name() + " ")
		fmt.Printf("\n")
	}
}
