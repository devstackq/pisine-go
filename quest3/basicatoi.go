package piscine

//"12345dasdsa"
//"12345"
//0

//1 convert num(string) string - num (int)
//2, +- not required
//
func BasicAtoi(s string) int {

	// string - rune array
	str := []rune(s) // array rune, []

	res := 0
	// set border - rune
	// range x=[48, 50,51...]
	for _, x := range str { // x 0: 49, 1: 55 ...
		if x >= '0' && x <= '9' { // [0] if rune -  number border >= 48 - <= 58 - border
			res = res*10 + int(x-48) // before  = rune, after - int // multiply*10 - (move digit) - 55 - 48 = int - num(7) - int method - do int
		}
	}

	// fmt.Print(res)

	return res
}

// // int multiply
// 	res = res * 10 + str[x]
// 	fmt.Print(res)
// }

// for range // before minus
// counter

//  n= n * - 1 // positive num
// // minus todo

// // PrintRune
