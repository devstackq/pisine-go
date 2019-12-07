package piscine

//010101010101

func BasicAtoi2(s string) int {
	// string - rune array // string - rune(ascii)
	str := []rune(s) // array rune, [0: 49, 1: 50, 2:51, 3:52,4:53] -
	res := 0         //save int

	// set border - rune
	// range x=[48, 50,51...]
	for _, x := range str { // x 0: 49, 1: 50 ... 49,

		if x >= '0' && x <= '9' { //  match - >='48' <= 57 50 true          [0] if rune -  number border >= 48 - <= 58 - border
			res = res*10 + int(x-48) // before  = rune, after - int // multiply*10 - (move digit) - 55 - 48 = int - num(7) - int method - do int
		} else {
			return 0
		}
	}
	return res // 12345
}

// 0 = 0 * 10 +  (49 - 48) // 1
// 1 = 1 * 10 + (50 - 48) // 10 + 2 = 12
// 12 = 12 * 10 + (51 - 48) / / 120 +3 = 123
