package main

import (
	"fmt"
	"strings"
)

func main() {
	files := WalkDirectory(".")

	ignoreRequest := []string{"main.go", "gitignore"}
	total := 0
	for k, v := range files {
		for i := 0; i < len(ignoreRequest); i++ {
			if strings.Contains(v.name, ignoreRequest[i]) {
				delete(files, k)
			}
		}
	}
	for _, v := range files {
		if v.lines != 0 {
			fmt.Println(v.name, v.lines)
		}
		total += v.lines
	}
	fmt.Println("Total lines in codebase:", total)
}
