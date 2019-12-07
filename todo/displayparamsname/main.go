package main

import (
	"os"

	"github.com/01-edu/z01"
)

// show last elem in array string

func main() {

	arg := os.Args // ./main hello there etc.., array strings[ 0main, 1hello, 1there...]

	count := 0 // count // geive all, or last elem inside args
	for range arg {
		count++
	}

	for _, value := range arg[count-1] { // show rune val, give last elem array string, [you]
		z01.PrintRune(value) //
	}
	z01.PrintRune('\n')
}
