package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// var files = make(map[string]FileInf)
var dir string
var files = make(map[string]FileInf)

func InteractCommand() map[string]FileInf {
	fmt.Println("Which directory's files would you like to count?")
	fmt.Println("OR type \"help\" to see a list of options")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	scanner.Scan()
	text := scanner.Text()
	cleaned := cleanInput(text)
	if len(cleaned) == 0 {
		fmt.Println("please enter in a directory name whose files you would like to count")
		os.Exit(1)
	}
	dir = cleaned[0]
	fmt.Println("Would you like to ignore any files? Press enter to skip this step")
	scanner.Scan()
	ignoreResponse := scanner.Text()
	if len(ignoreResponse) == 0 {
		return WalkDirectory(dir)
	}
	cleanedFilesIgnored := cleanInput(ignoreResponse)
	files = WalkDirectory(dir)
	fmt.Printf("- Counting files for directory: %s\n- Ignoring files: %v\n", dir, cleanedFilesIgnored)
	fmt.Println("- - - - - - - - - - - - - - - - - - - ")
	time.Sleep(time.Second * 1)
	dirWithIgnoredFiles := IgnoreRequestFunc(cleanedFilesIgnored, files)
	return dirWithIgnoredFiles
}

// IgnoreRequestFunc takes a string slice of file names which have been requested to be ignored in the count. The slice is iterated over to get its values which are checked against the values of the files map, in this case, the value v.Name of the FileInf struct.
// If these match, the file is removed from the files map and is returned to be counted in the rest of the program flow.
func IgnoreRequestFunc(cleanedFilesIgnored []string, files map[string]FileInf) map[string]FileInf {
	for i := 0; i < len(cleanedFilesIgnored); i++ {
		for k, v := range files {
			if strings.Contains(v.Name, cleanedFilesIgnored[i]) {
				delete(files, k)
			}
		}
	}
	return files
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
