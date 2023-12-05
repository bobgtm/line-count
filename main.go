package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	files := WalkDirectory(".")

	excludeFile := flag.Bool("f", false, "Name files to exclude from count")
	flag.Parse()

	total := 0
	switch {
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
