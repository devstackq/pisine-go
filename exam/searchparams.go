package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]
	if len(args) == 3 {
		runa := []rune(args[0]) // Pipe //  "Hello!" // // [72 101 108 111 33]

		runa2 := []rune(args[1])[0] // replace
		// fmt.Printf("%v", runeVal)

		runa3 := []rune(args[2])[0]

		for idx, v := range runa { //         runeVal  - length
			if v == runa2 { // match 3-1=2 == 2,
				runa[idx] = runa3 // return [2:76] l
			}
		}
		for _, v := range runa {
			z01.PrintRune(v)

		}
		z01.PrintRune(10)
	} else {
		z01.PrintRune(10)
	}
}
