package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	//
	args := os.Args[1:]

	l := len(args)
	if l == 3 {

		runes := []rune(args[0])      // 1 word
		word := []rune(args[1])[0]    // replace letter
		replace := []rune(args[2])[0] // letter

		for idx, v := range runes {
			if v == word {
				runes[idx] = replace
			}

		}
		for _, v := range runes {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('\n')
		return
	}
}
