package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	// name := ""        // string create
	name := os.Args[0] // os args [0]

	for _, x := range name { //. / m / a/ i / n

		z01.PrintRune(x)
	}
	z01.PrintRune(10)
}
