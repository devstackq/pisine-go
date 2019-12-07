package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {

	//table  count 4
	//[0HOw 1are 2 you?]
	count := 0
	str := "" //How are you? friend
	nl := "\n"
	//4
	for range table {
		count++
	}

	for i := 0; i < count; i++ { //0<3, 1<3,2<3,
		str = str + table[i] // "" = ""+ How| How = How + are + you?
		if count-1 > i {     // 3>0, newlnine,  3>1, newline 3>2, 3>3 false
			str = str + nl
		}
	}
	//convert to rune and print value
	runeVal := []rune(str)
	for _, v := range runeVal {
		z01.PrintRune(v)

	}
	z01.PrintRune(10)
}
