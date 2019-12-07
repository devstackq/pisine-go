package piscine

func Swap(a *int, b *int) {
	var c int
	c = *a
	*a = *b
	*b = c

}

// a = a + b // 17
// b = a -b 17 - 8 = 9
// a = a -b 17 - - 9 = 8

// a = a-b
// b = b + a
// a = b - a
