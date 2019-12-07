package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hedasdasdasdallo!", 5))

}

//return nth rune, else return 0
//hello!, n =3
func NRune(s string, n int) rune {

	runeVal := []rune(s) //  "Hello!" // // [72 101 108 111 33]

	// fmt.Printf("%v", runeVal)

	for ikldjas := range runeVal { //         runeVal  - length
		fmt.Print(idx) 0,1,2,3,4,56
		fmt.Print(value)
		if idx == n-1 { // match 3-1=2 == 2,
			return runeVal[idx] // return [2:76] l
		}
	}
	return 0 // all case return 0
}

//count
// index +1
// comapre index == n -1

// x = 0
// y = 3

// x0 y3 5 6

// [y            ] [x          ]

// os.Args[]
func NRune(s []string, r, w string) string {

	runeVal := []rune(s) //  "Hello!" // // [72 101 108 111 33]

	// fmt.Printf("%v", runeVal)

	for idx, v := range runeVal { //         runeVal  - length
		
		
		if v == w { // match 3-1=2 == 2,
			return runeVal[idx] = r // return [2:76] l
		}
	}
	return runeVal
}
