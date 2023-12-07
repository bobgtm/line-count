package main

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

func interactCommand() map[string]FileInf {
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
	dirWithIgnoredFiles := ignoreRequestFunc(files, cleanedFilesIgnored)
	return dirWithIgnoredFiles
}

func ignoreRequestFunc(files map[string]FileInf, cleanedFilesIgnored []string) map[string]FileInf {
	for k, v := range files {
		for i := 0; i < len(cleanedFilesIgnored); i++ {
			if strings.Contains(v.name, cleanedFilesIgnored[i]) {
				delete(files, k)
			}
		}
	}
	return files
}
