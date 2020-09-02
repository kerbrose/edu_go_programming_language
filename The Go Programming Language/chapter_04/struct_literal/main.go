// testing struct import
package main

import (
	"fmt"

	"./p"
)

func main() {
	pp := p.U{A: 3, B: 4}
	uu := p.U{3, 5, 0}
	fmt.Println(pp)
	fmt.Println(uu)
	//var _ = p.T{a: 1, b: 2} // compile error: can'treference a, b
	//var _ = p.T{1, 2}   // compile error: can'treference a, b
}
