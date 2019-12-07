package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(39)
	PrintNbr(0)
	PrintNbr(-192)
}

func PrintNbr(n int) { // 4627

if n < 0 {
	z01.PrintRune('0')
}

PrintNum(n)
}

PrintNum(num int){
	x :='0'
	if num ==0 {
		z01.PrintRune('0')
		return
	}

	for a:=1; a <= num%10; a++ {  //9, a2
		for b:= -1; b >= num%10; b-- {//9
			x++
		}
		if num/10 !=0 { //3
			PrintNum(num /10)
		}
	}
	z01.PrintRune(x)
	return
}








	// var negVal boo
	// if n < 0 {
	// 	z01.PrintRune('-')
	// }

	// if n == -9223372036854775808 {
	// 	n = -9223372036854775808
	// 	n = n * -1
	// 	negVal = true
	// } else {
	// 	n = n * -1
	// 	z01.PrintRune(rune(n))
	// }

// 	if n > 0 {

// 		arrNums := []int{}
// 		count := 0
// 		for range arrNums {
// 			count++
// 		}

// 		if n == 0 {
// 			z01.PrintRune('0')
// 		}

// 		if n > 0 {
// 			for i := 0; i <= 100; i++ {
// 				mod := n % 10
// 				arrNums = append(arrNums, mod)
// 				n = n / 10

// 			}
// 			fmt.Print(arrNums)

// 			for i := 0; i <= count; i++ {
// 				z01.PrintRune(rune(arrNums[i]))
// 			}
// 		}
// 	}
// }

// // if negVal { //	for range arrNums {
// 			count++
// 		} - 9223372036854775808,
// // 	{
// // 		z01.Print	for range arrNums {
// 			count++
// 		}'8') // print 8 - backspace ?
// // 	}
// // }
// } else {
// 	z01.PrintRune('0') // else return 0
// }


