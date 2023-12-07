package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// If -n is used, program runs in non-interactive mode
	interactive := flag.Bool("i", false, "Run Program in interactive mode")
	excludeFile := flag.Bool("f", false, "Name files to exclude from count")
	flag.Parse()
	var files = make(map[string]FileInf)
	var dir string
	total := 0
	switch {
	case *interactive:
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
		fmt.Println("Which files would you like to ignore?")
		scanner.Scan()
		ignoreRequest := scanner.Text()
		cleanedFiles := cleanInput(ignoreRequest)
		files = WalkDirectory(dir)
		for k, v := range files {
			for i := 0; i < len(cleanedFiles); i++ {
				if strings.Contains(v.name, cleanedFiles[i]) {
					delete(files, k)
				}
			}
		}

		fmt.Printf("Counting files for directory: %s\n", dir)
		time.Sleep(time.Second * 3)
	case *excludeFile:
		ignoreRequest := os.Args[2:]
		for k, v := range files {
			for i := 0; i < len(ignoreRequest); i++ {
				if strings.Contains(v.name, ignoreRequest[i]) {
					delete(files, k)
				}
			}
		}
	default:
		break
	}

	for _, v := range files {
		if v.lines != 0 {
			fmt.Println(v.name, v.lines)
		}
		total += v.lines
	}
	fmt.Println("Total lines in codebase:", total)
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
