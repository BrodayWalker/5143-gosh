package main

import(
    "fmt"
    "os"
    //"path/filepath"
)

// A command grouped with arguments for calling it
type CommandLine struct {
    comm string
    args []string
}

// PipeLine
// Special version of the execute function which takes a list of Commands, then
// for each Command, sends it's output to a file and uses that file as the first
// arg of the next command.
func PipeLine(commands []CommandLine){

    fmt.Println("PipeLine starting!\nfmt imported for debugging...")

    // stdout backup
    stdout := os.Stdout;
    // Path to the pipe file
    //pipeFilePath := filepath.Join(os.TempDir(), "gosh.pipe.tmp")
    pipeFilePath := "gosh.pipe.tmp"
    // Create the actual pipe file
    pipeFile, _ := os.Create(pipeFilePath)
    pipeFile.Close()

    // For each command in the array
    for i, pipe := range commands {

        // DEBUGGING: Print the command and args set to execute
        fmt.Println("\n\nSetting up to Execute...")
        fmt.Println("Command: ", pipe.comm)

        // If this isn't the last command
        if i < len(commands) - 1 {
            // Before processing each command, open the file and redirect stdout
            os.Stdout, _ = os.Open(pipeFilePath)
        }

        // If the Command is valid
        if com, valid := ComMap[pipe.comm]; valid{         
            // If this isn't the first command
            if i > 0{
                // We need to add the pipe file to the args (at the front)
                pipe.args = frAddStr(pipe.args, pipeFilePath)
            }

            // DEBUGGING: Print the updated arguments and what the output SHOULD be (w/o  file redirection)
            backupOut := os.Stdout
            os.Stdout = stdout
            fmt.Println("Args: ", pipe.args)
            fmt.Println("Output: ")
            com(pipe.args)
            os.Stdout = backupOut

            // Execute the command with its arguments
            com(pipe.args)
        }
        
        // If this isn't the last command
        if i < len(commands) - 1 {
            // After processing each command, close the pipe file
            err := os.Stdout.Close()
        }

        // After processing each command, restore stdout
        os.Stdout = stdout

        // DEBUGGING: print the current contents of temp file after each command
        fmt.Println("\nSTATUS OF TEMP FILE\n##############################################")
        debugFile, _ := os.Open(pipeFilePath)
        fmt.Println(debugFile)
        debugFile.Close()

    }

}

func frAddStr(argList []string, arg string) []string {
    argList = append(argList, "")
    copy(argList[1:], argList)
    argList[0] = arg
    return argList
}
