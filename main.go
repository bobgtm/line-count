package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// If -n is used, program runs in non-interactive mode
	interactive := flag.Bool("i", false, "Run Program in interactive mode")
	excludeFile := flag.Bool("f", false, "Name files to exclude from count")
	flag.Parse()
	total := 0
	switch {
	case *interactive:
		interactCommand()
	case *excludeFile:
		ignoreRequest := os.Args[2:]
		ignoreRequestFunc(files, ignoreRequest)
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
