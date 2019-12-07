package main

import (
	"os"

	"github.com/01-edu/z01"
)

//write arg, command line
// ./ main gogo dance cat
func main() {

	count := 0

	//[0: ./main, 1: text, 2: gogo, 3:cat]
	a := os.Args[1:]
	//save var a - data -terminal - array string
	//counter
	for range a {
		count++ //4
	}
	// if i==0, new line
	//i ==0, k return empty line
	//i==1, k, from start to end index a[1], and printRune(k)

	for i := 0; i < count; i++ { // 1
		for _, k := range a[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}
