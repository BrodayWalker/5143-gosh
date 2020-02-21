package main

// A command grouped with arguments for calling it
type CommandLine struct {
    comm Command
    args []string
}

// PipeLine
// Special version of the execute function which takes a list of Commands, then
// for each Command, sends it's output to a file and uses that file as the first
// arg of the next command.
func PipeLine(commands []CommandLine){

    // For each command in the array
    for i, pipe := range commands {

        // If the Command is valid
        if com, valid := ComMap[pipe.comm]; valid{
            
            // If this isn't the first command
            if i > 0{
                // We need to add the pipe file to the args (at the front)
                frAddStr(pipe.args, "gosh.pipe.tmp")
            }

            // If the command isn't the last command in the PipeLine
            // Send its output to the pipe file
            // Else (it is the last command)
            // Restore stdout

            // Execute the command with its arguments
            com(pipe.args)
        }

    }

    // Route the command to call the proper function
    if com, valid := ComMap[command]; valid{
        com(args)
    }else if command == "exit"{
        os.Exit(0)
    }else{
        fmt.Println("Command not found.")
    }
}

func frAddStr(argList []string, arg string) []int {
    argList = append(argList, 0)
    copy(argList[1:], argList)
    argList[0] = arg
    return argList
}