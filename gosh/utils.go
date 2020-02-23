package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ValidPath checks the validity of a path
func ValidPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	return false
}

// BuildPathToFile accepts a string path to a file and creates an absolute
// path to that file.
func BuildPathToFile(path string) string {
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

// FileExists checks to see if a file exists at a specified path.
// If the file exists, FileExists returns true
func FileExists(file string) bool {
	fInfo, err := os.Stat(file)
	// Will evaluate to true if the error is because the file does
	// not exist
	if os.IsNotExist(err) {
		return false
	}
	// Will return false (file does not exist) if the string argument
	// is a directory
	// Will return true if the file exists and is not a directory
	return !fInfo.IsDir()
}

// ArgSplitter separates flags from other arguments and explodes them
// Regular arguments are returned in an array of strings
// Flags are returned in exploded for as an array of strings
// Example case: -lha myFile.txt -b -x file2.txt
// returns [myFile.txt file2.txt] [l h a b x]
func ArgSplitter(args []string) (argList []string, flags []string) {
	// Compile regexp object to test against
	var flagMatch, _ = regexp.Compile(`^-{1,2}(\D)+`)

	if len(args) != 0 {
		// Iterate through arguments and split
		for i := range args {
			// Evaluates to true if one or more flags are present
			if flagMatch.MatchString(args[i]) {
				// Split flags
				split := strings.SplitN(args[i], "", len(args[i]))
				// Push flags onto flag slice
				for j := range split {
					// Only push letters onto slice
					if split[j] != "-" {
						flags = append(flags, split[j])
					}
				}

			} else {
				// args[i] is not a flag; push argument on argList
				// This is a stable operation.
				argList = append(argList, args[i])
			}
		}
	} else {
		fmt.Println("No args to split")
	}

	return argList, flags
}
