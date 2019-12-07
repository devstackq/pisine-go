package piscine

func TrimAtoi(s string) int {
	//"sd - x1fa2W3s4"
	res := 0
	n := 1                //, change -1, if match (-) -> 14 line
	for _, x := range s { // break up rune each symbol, 54,23,55,12,55,22...
		if x >= '0' && x <= '9' { // if match, from 48 - 57 do 9 line
			res = res*10 + int(x-48) // 1*10 + int(49 -1) = 1
		} else if res == 0 && x == '-' { // res =0, and change value - line 6,
			n = -1 //  n = -1, after line 15,
		}
	}
	return res * n // done result * -1, = 12345 * -1 = -12345
}



package main

import (
	"fmt"
)

func main() {
	fmt.Println(Index("Hello!", "lp"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))

}

func Index(s string, toFind string) int {

	//counter 1

	var lenT, lenF, counter int

	//count1
	for range s {
		lenT++
	}

	//
	for range toFind {
		lenF++
	}
	//toFine =3,
	//s = 10
	//match s == toFine
	//return index
	for i := 0; i <= lenT-lenF; i++ {
		for j := 0; j < lenF; j++ {
			if s[i] == toFind[0] {

				counter++ //3
			}
			if counter == lenF {
				return i
			}
		}
	}
	return -1

}
