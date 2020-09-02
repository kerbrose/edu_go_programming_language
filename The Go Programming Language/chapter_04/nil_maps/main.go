// Testing the difference between nil maps & empty maps
package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages == nil)    // "false"
	fmt.Println(len(ages) == 0) // "false"
	var ages2 map[string]int
	//panic: assignment to entry in nil map
	ages2["carol"] = 21
	// the panic could be avoided if initialized
	ages2 = make(map[string]int)
	fmt.Println(ages2 == nil)    // "true"
	fmt.Println(len(ages2) == 0) // "true"
}
