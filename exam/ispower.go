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
	f := "false" //string - >

	if content, err := strconv.Atoi(args[0]); err == nil {

		// conent = 7
		for content > 1 {
			if content%2 == 1 {
				for _, v := range f {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
				return
			}
			content = content / 2 // 256/2 = 128, 64, 32, 16, 8,4,2,1
		}
		for _, v := range t {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	}
	}
