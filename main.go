package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pathString := os.Getenv("PATH")
	paths := strings.Split(pathString, ":")

	pathCounts := make(map[string]int)
	binsPaths := make(map[string][]string)

	for _, path := range paths {
		pathCounts[path]++
	}

	for path := range pathCounts {
		binaries, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}

		for _, binary := range binaries {
			name := binary.Name()

			slice, found := binsPaths[name]
			if !found {
				slice = []string{path}
			} else {
				slice = append(slice, path)
			}
			binsPaths[name] = slice
		}
	}

	for path, count := range pathCounts {
		if count <= 1 {
			continue
		}

		fmt.Printf("Path %s is defined %d times on the PATH\n", path, count)
	}

	for binary, binPaths := range binsPaths {
		if len(binPaths) <= 1 {
			continue
		}

		fmt.Printf("%s\n", binary)
		for _, path := range binPaths {
			fmt.Printf("\t%s\n", path)
		}
	}
}
