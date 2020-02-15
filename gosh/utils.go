package main

import (
	"os"
)

// ValidPath returns a boolean value if the file already exists
func ValidPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	return false
}
