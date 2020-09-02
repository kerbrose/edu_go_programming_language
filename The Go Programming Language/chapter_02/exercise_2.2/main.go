/*
Exercise 2.2: Write a general-purpose unit-conversion program analogous
to cf that reads numbers from its command-line arguments or from the standard
input if there are no arguments, and converts each number into units like
temperature in Celsius and Fahrenheit, length in feet and meters, weight
in pounds and kilograms, and the like.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"./lengthconv"
	"./tempconv"
	"./weightconv"
)

var mType = flag.String("t", "", "Measurement Type: [l: Length, t: Temperature, w: Weight]")

func main() {
	flag.Parse()
	switch *mType {
	case "l":
		for _, arg := range flag.Args() {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f := lengthconv.Feet(v)
			m := lengthconv.Meter(v)
			fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), m, lengthconv.MToF(m))
		}
	case "t":
		for _, arg := range flag.Args() {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			c := tempconv.Celsius(v)
			f := tempconv.Fahrenheit(v)
			k := tempconv.Kelvin(v)
			fmt.Printf(
				"%s = %s = %s, %s = %s = %s, %s = %s = %s\n",
				c, tempconv.CToF(c), tempconv.CToK(c),
				f, tempconv.FToC(f), tempconv.FToK(f),
				k, tempconv.KToC(k), tempconv.KToF(k))
		}
	case "w":
		for _, arg := range flag.Args() {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			k := weightconv.Kilogram(v)
			p := weightconv.Pound(v)
			fmt.Printf(
				"%s = %s, %s = %s\n",
				k, weightconv.KToP(k),
				p, weightconv.PToK(p))
		}
	default:
		fmt.Println(flag.Args())
	}
}
