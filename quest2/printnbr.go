package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) { // 4627

	var negVal bool
	if n < 0 {
		z01.PrintRune('-')
	}

	if n == -9223372036854775808 {
		n = -9223372036854775808
		n = n * -1
		negVal = true
	} else {
		n = n * -1
		z01.PrintRune(rune(n))
	}

	if n > 0 {

		count := 0
		text := ""

		if n == 0 {
			z01.PrintRune('0')
		}
		for i := 0; i <= n; i++ {
			if n > 0 {
				n := n % 10                // 562 = 2, 6
				text = text + string(n+48) //7, ''+ 2+48, 50, ''+string(50) = 2, string48+6= "2"+ "6", "5" - 87434565     265
				n = n / 10                 // 562/10, 56, n56
				count++                    //1
			}
			for _, v := range text {
				z01.PrintRune(v)
			}

		}

		if negVal { // true - 9223372036854775808,
			{
				z01.PrintRune('8') // print 8 - backspace ?
			}
		}
	} else {
		z01.PrintRune('0') // else return 0
	}

}
