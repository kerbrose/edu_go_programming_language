// Exercise 1.3: Experiment to measure the difference in running
// time between our potentially inefficient versions and the one
// that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests
// for systematic performance evaluation.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	BenchmarkEcho1()
	BenchmarkEcho2()
}

func BenchmarkEcho1() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func BenchmarkEcho2() {
	start := time.Now()
	s := strings.Join(os.Args[:], " ")
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
