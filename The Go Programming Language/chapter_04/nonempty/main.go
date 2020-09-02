// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	data2 := []string{"one", "", "three", "", "five"}
	fmt.Printf("%q\n", data)
	fmt.Printf("%q\n", nonempty(data)) //`["one" "three"]`
	fmt.Printf("%q\n", data)           //`["one" "three" "three"]`
	fmt.Printf("%q\n", nonempty(data2))
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
