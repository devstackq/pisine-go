package piscine

func Atoi(s string) int {
	minus := false
	positive := false

	// string - rune array // string - rune(ascii)
	str := []rune(s) // array rune, [0: 49, 1: 50, 2:51, 3:52,4:53] -
	res := 0         //save int

	// set border - rune
	// range x=[48, 50,51...]

	for i, x := range str { // x 0: 49, 1: 50 ... 49,
		//-453

		if x == 45 && i == 0 { // -
			minus = true
			//+432
		} else if x == 43 && i == 0 { // +
			positive = true
		} else if x >= '0' && x <= '9' {
			//  match - >='48' <= 57 50 true          [0] if rune -  number border >= 48 - <= 58 - border
			res = res*10 + int(x-48) // before  = rune, after - int // multiply*10 - (move digit) - 55 - 48 = int - num(7) - int method - do int // 1234
		} else {
			return 0
		}
	}
	if minus {
		res *= -1 // 1234 * - 1 = -1234
	}
	// -
	if positive { //
		res *= 1 // 1234 * 1 = 1234
	}

	return res
}
