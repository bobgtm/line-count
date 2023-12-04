package main

import (
	"fmt"
)

func main() {

	files := WalkDirectory(".")

	for _, v := range files {
		if v.lines != 0 {
			fmt.Printf("%s: %d line(s)\n", v.name, v.lines)
		}

	}

}
