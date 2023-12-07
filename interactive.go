package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// var files = make(map[string]FileInf)
var dir = "."
var files = WalkDirectory(dir)

func interactCommand() {
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
	ignoreRequest := scanner.Text()
	if len(ignoreRequest) == 0 {
		return
	}
	cleanedFiles := cleanInput(ignoreRequest)

	ignoreRequestFunc(files, cleanedFiles)
	fmt.Printf("- Counting files for directory: %s\n- Ignoring files: %v\n", dir, cleanedFiles)
	fmt.Println("- - - - - - - - - - - -")
	time.Sleep(time.Second * 3)
}

func ignoreRequestFunc(files map[string]FileInf, ignore []string) map[string]FileInf {
	for k, v := range files {
		for i := 0; i < len(ignore); i++ {
			if strings.Contains(v.name, ignore[i]) {
				delete(files, k)
			}
		}
	}
	return files
}
