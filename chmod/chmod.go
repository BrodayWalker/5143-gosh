package main

//For this command, the function will receive (hopefully) the portion of the command
//after 'chmod '. It will process the options by splitting the arguments string
//and then iteratively traversing each substring.

//TO COMPILE AND EXECUTE THE CODE: go build; ./chmod
//Make sure to navigate to the chmod directory before building and executing

//**************************************************************************//
//NEED TO FIND A WAY TO MODIFY PERMISSION BITS
//**************************************************************************//

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//chmod u=rwx,g=rx,o=r myfile

func main() {
	initial := "577 ooga.txt"
	input := strings.Fields(initial) //split string apart by whitespaces
	// fmt.Println(string(input[0]))
	// fmt.Println(string(input[2]))
	// fmt.Println(string(input[7]))
	// fmt.Println(strings.Fields(input)[1])

	os.OpenFile("."+input[len(input)-1], os.O_CREATE, 0000)

	test1, err := os.Lstat("C:\\Users\\Owner\\Desktop\\5143-OS-Matamoros\\5143-gosh\\chmod\\.ooga.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("permissions: %#o\n", test1.Mode().Perm()) //https://golang.org/pkg/os/?m=all#fileStat
	// mode := int(0666)
	// mode2 := os.FileMode(mode)
	// for _, argument := range input {
	// 	fmt.Println("This is the argument: ", string(argument))
	// 	for _, character := range string(argument) {
	// 		fmt.Println(string(character))
	// 	}
	// }

	//easy part 1: if someone just typed 'chmod' with no arguments

}

// func CopyFile(src, dst string) (err error) {

// }
