package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//path := ""
	err := os.Mkdir("test", 755)

	fmt.Printf("%T", err)

	if err != nil {
		log.Fatal(err)
	}
}
