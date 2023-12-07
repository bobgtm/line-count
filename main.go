package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bobgtm/linecount/internal"
)

func main() {

	// If -n is used, program runs in non-interactive mode
	interactive := flag.Bool("i", false, "Run Program in interactive mode")
	excludeFile := flag.Bool("f", false, "Name files to exclude from count")
	flag.Parse()
	total := 0
	files := internal.WalkDirectory(".")
	switch {
	case *interactive:
		files = internal.InteractCommand()
	case *excludeFile:
		ignoreRequest := os.Args[2:]
		files = internal.IgnoreRequestFunc(ignoreRequest, files)
	default:
		break
	}

	for _, v := range files {
		if v.Lines != 0 {
			fmt.Println(v.Name, v.Lines)
		}
		total += v.Lines
	}
	fmt.Println("Total lines in codebase:", total)
}
