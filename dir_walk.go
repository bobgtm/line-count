package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileInf struct {
	name  string
	lines int
}

func WalkDirectory(dir string) map[string]FileInf {
	// Initialize empty map to hold file name path and the line count
	fileLineCount := make(map[string]FileInf)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("was unable to access directory path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "git") {
			return filepath.SkipDir
		}

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		reader := bufio.NewReader(file)
		lineCount := 0
		for {
			_, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			lineCount++
		}

		files := FileInf{
			file.Name(),
			lineCount,
		}
		fileLineCount[path] = files

		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return fileLineCount
}
