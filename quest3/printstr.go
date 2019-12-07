package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	text := []rune(str)
	for i := range text {
		z01.PrintRune(text[i])
	}
}
