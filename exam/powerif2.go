package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]

	count := 0
	for range args {
		count++
	}

	if count != 1 {
		z01.PrintRune('\n')
		return
	}

	t := "true"
	f := "false"

	if i, err := strconv.Atoi(args[0]); err == nil {
		for i != 1 {
			if i%2 == 1 {
				for _, v := range f {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return
			}
			i = i / 2
		}

		for _, v := range t {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		return
	}

}
