package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	count := 0
	a := os.Args

	for range a {
		count++
	}

	for i := count; i > 1; i-- {

		for _, k := range a[i-1] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}

}
