package main

import(
    "io"
    "io/ioutil"
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
    
    // stdout backup
    stdout := os.Stdout;
    // Path to the pipe file
    //outPipePath := filepath.Join(os.TempDir(), "gosh.pipe.tmp")
    dir, err := os.Getwd()
    outPipeName := "gosh.pipe.out.tmp"
    if err != nil {
        fmt.Println("Error getting working directory: ", err)
        return
    }
    outPipePath := dir + "/" + outPipeName
    inPipeName := "gosh.pipe.in.tmp"
    if err != nil {
        fmt.Println("Error getting working directory: ", err)
        return
    }
    inPipePath := dir + "/" + inPipeName

    // For each command in the array
    for i, pipe := range commands {

        // DEBUGGING: Print the command and args set to execute
        fmt.Println("\n\nSetting up to Execute...")
        fmt.Println("Command: ", pipe.comm)

        // If this isn't the last command
        if i < len(commands) - 1 {
            // Before processing each command, open the file and redirect stdout
            os.Stdout, err = os.OpenFile(outPipePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
            if err != nil {
                os.Stdout = stdout
                fmt.Println("Error opening temp file: ", err)
                return
            }
        }

        // If the Command is valid
        if com, valid := ComMap[pipe.comm]; valid{         
            // If this isn't the first command
            if i > 0{
                // We need to add the pipe file to the args (at the front)
                pipe.args = frAddStr(pipe.args, inPipeName)
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
            // Move file's pointer back to start of file
			os.Stdout.Seek(0, io.SeekStart)
        }
        
        // If this isn't the last command
        if i < len(commands) - 1 {
            // After processing each command, close the pipe file
            err = os.Stdout.Close()
            if err != nil {
                fmt.Println("Error closing temp file: ", err)
                return
            }
        }

        // After processing each command, restore stdout
        os.Stdout = stdout

        // Copy the output to the input file for the next command
        copyFrom(outPipePath, inPipePath)

        // DEBUGGING: print the current contents of temp file after each command
        fmt.Println("\nSTATUS OF TEMP FILE\n##############################################")
        contents, err := ioutil.ReadFile(outPipeName)
        if err != nil {
            os.Stdout = stdout
            fmt.Println("Error opening temp file (for status check): ", err)
            return
        }
        fmt.Println(string(contents))

    }

}

func frAddStr(argList []string, arg string) []string {
    argList = append(argList, "")
    copy(argList[1:], argList)
    argList[0] = arg
    return argList
}

// File copy borrowed from
// https://opensource.com/article/18/6/copying-files-go
func copyFrom(src, dst string) (int64, error) {
        sourceFileStat, err := os.Stat(src)
        if err != nil {
                return 0, err
        }

        if !sourceFileStat.Mode().IsRegular() {
                return 0, fmt.Errorf("%s is not a regular file", src)
        }

        source, err := os.Open(src)
        if err != nil {
                return 0, err
        }
        defer source.Close()

        destination, err := os.Create(dst)
        if err != nil {
                return 0, err
        }
        defer destination.Close()
        nBytes, err := io.Copy(destination, source)
        return nBytes, err
}