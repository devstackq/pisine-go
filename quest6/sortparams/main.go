package main

import (
	"os"

	"github.com/01-edu/z01"
)

//array string -> string -> runes -> sort bubble
func main() {
	arg := os.Args[1:] //[ 0: ./main, 1: H, 2:W, 3: 3...] // arrays string

	countArg := 0     // 8
	countRune := 0    // 8
	var strVal string //

	for range arg {
		countArg++
	}

	//count string - os>args // [strings] - |a|b|c|2|r|6|a..
	for i := 0; i < countArg; i++ { //8
		strVal = strVal + arg[i] // "" + string // 1loop "" + D, 2loop D+f.... DF3lkj423d
	}

	//converr  string letter - to rune
	//[0:D -> 87, 1: f -> 77, 2: 3 - 51......]
	runs := []rune(strVal) //

	for range runs {
		countRune++
	}
	//8

	// 0: D - 99 > 1: f - 69
	// 69, 99
	// sort rune, ascii and order
	for i := 0; i < countRune-1; i++ { // 0 | 35,
		for j := i; j < countRune; j++ { // j 0 | 36
			if runs[i] > runs[j] {
				runs[i], runs[j] = runs[j], runs[i]
			}
		}
	}
	//runs - sorted 1134DOdasds...

	for i := 0; i < countRune; i++ { //8
		z01.PrintRune(runs[i]) // 1, 1,3....
		z01.PrintRune(10)
	}
}

// func main() {

// 	arg := os.Args

// 	count := 0

// 	for range arg {
// 		count++
// 	}

// 	for i := 0; i < count-1; i++ { //
// 		for j := i; j < count; j++ {
// 			for _, x := range arg {
// 				rv := []rune(x)
// 				if rv[i] > rv[j] {
// 					rv[i], rv[j] = rv[j], rv[i]
// 				}
// 				z01.PrintRune(rv[i])
// 			}

// 		}
// 	}

// }
