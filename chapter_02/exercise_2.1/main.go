/*
Exercise 2.1: Add types, constants, and functions to tempconv for processing
temperatures in the Kelvin scale, where zero Kelvin is −273.15°C and a difference
of 1K has the same magnitude as 1°C.
*/

//using package tempconv
package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Printf("Cold! %v\n", tempconv.BoilingK)
	fmt.Printf("Kelvin! %v\n", tempconv.CToK(0))
	k1 := tempconv.Kelvin(273)
	fmt.Println(tempconv.KToC(k1))
	fmt.Println(tempconv.KToF(k1))
}
