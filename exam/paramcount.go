package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]
	if len(args) == 3 {
		runa := []rune(args[0]) // Pipe //

		runa2 := []rune(args[1])[0] // word

		runa3 := []rune(args[2])[0] // replace

		for idx, v := range runa { //
			if v == runa2 { // match 3-1=2 == 2,
				runa[idx] = runa3 // return [2:76] l
			}
		}
		for _, v := range runa {
			z01.PrintRune(v)

		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
