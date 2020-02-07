package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := ".\\test"
	var mode os.FileMode = 0755

	err := os.Mkdir(path, mode)

	fmt.Printf("%T", err)

	if err != nil {
		log.Fatal(err)
	}
}
