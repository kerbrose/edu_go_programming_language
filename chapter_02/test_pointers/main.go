// checking the difference between python & golang
package main

import (
	"fmt"
)

func main() {
	var x int = 3
	var y int = 3
	// checking if pointers hold the same address, in python this return true
	fmt.Println((&x == &y))
}
