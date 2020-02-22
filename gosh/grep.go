package main

import (
	"os"
)

func init(){
    // Add this command's function to the command mapping
    ComMap["grep"] = Grep
}

func Grep(args []string){

}
