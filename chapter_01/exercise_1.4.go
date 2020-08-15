// Exercise 1.4: Modify dup2 to print the names of all files in which
// each duplicated line occurs.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileLines := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		fileLines[filename] = make(map[string]int)
		for _, line := range strings.Split(string(data), "\n") {
			fileLines[filename][line]++
		}
	}

	for filename, lines := range fileLines {
		for line, n := range lines {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}
