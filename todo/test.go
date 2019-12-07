package piscine

import "github.com/01-edu/z01"

// import (
// 	"github.com/01-edu/z01"
// )

// func IterativePower(nb int, power int) int {
// 	if power == 0 {
// 		return 1
// 	}
// 	if power < 0 {
// 		return 0
// 	}
// 	if power == 1 {
// 		return nb
// 	}

// 	valnb := nb //
// 	for i := 1; i <= power; i++ {
// 		nb = valnb * nb
// 	}
// 	return nb
// }

// func RecursivePower(nb int, power int) int {

// 	if nb == 0 {
// 		return 1
// 	}
// 	if nb < 0 {
// 		return 0
// 	}
// 	if nb == 1 {

// 	}
// 	return nb * RecursivePower(nb-1, power)

// }

// func strrev(s string) string {
// 	new := []rune(s)
// 	old := []rune(s)
// 	count := 0

// 	for range old {
// 		count++
// 	}
// 	//5
// 	for i := 0; i < count; i++ { // 1<5
// 		new[count-i-1] = old[i] // 5-0-1, new 4 = old[0]
// 	}
// 	return string(new)
// }

// func RecursivePower(nb int, power int) int {
// 	if power == 0 {
// 		return 1
// 	}
// 	if power < 0 {
// 		return 1
// 	}

// 	return nb * RecursivePower(nb-1, power)
// }

// func Fibonacci(index int) int {
// 	var result int
// 	if index < 0 {
// 		return -1
// 	}
// 	if index == 0 || index == 1 {
// 		return index
// 	}

// 	result = Fibonacci(index-1) + Fibonacci(index-2)

// 	return result
// }

// package piscine

// func Sqrt(nb int) int {

// 	var b int
// 	var c int

// 	for a := 1; a <= nb; a++ { // 25, 5 loop
// 		if a > 0 {
// 			b = nb / a               // b = 25/5 =5
// 			if nb%a == 0 && b == a { // 25/5 == 0 T, 5 == 5 T
// 				return b // 5
// 			}
// 		}
// 	}
// 	return c
// }

// func AlphaCount(str string) int {

// 	count:=0;

// 	for _, x:= range str { //
// 		if x >= 65 && x <= 90 || x >= 97 && x <= 122 {
// 			count++//
// 		}

// 	}
// 	return count
// }

// //

func FirstRune(s string) rune {

	str := []rune(s)
	var x int
	for x := range str { // str[0:1] - H
		if x >= 'A' && x <= 'Z' {
			return str[x] // [0:H]
		}
	}
	return str[x] //
}

// str:= []rune(s)

// z01.PrintRune(s[0])

func NRune(s string, n int) rune {
	str := []rune(s)

	for x := range str {
		if x == n-1 {
			z01.PrintRune(str[x])
		}
	}
	return 0
}

func LastRune(s string) rune {
	str := []rune(s)

	count := 0
	for range str {
		count++
	}
	z01.PrintRune(str[count-1])

	return 0

}

func Compare(a, b string) int {

	if a == b {
		return 0
	}
	if a < b {
		return -1
	} else {
		return 1
	}
}

func ToUpper(s string) string {

	str := []rune(s)

	for x := range str {
		if str[x] >= 97 && str[x] <= 122 { // comapare value,
			str[x] = str[x] - 32 // change - to upper, -> index, where have low letter
		}
	}
	return string(str)
}

func ToLower(s string) string {

	str := []rune(s)

	for x := range str {
		if str[x] >= 65 && str[x] <= 90 { // comapare value,
			str[x] = str[x] + 32 // change - to upper, -> index, where have low letter
		}
	}
	return string(str)
}

func IsAlpha(str string) bool {

	strA := []rune(str)
	flag := true

	for _, x := range strA {
		// if x >= 65 && x <= 90 || x >= 97 && x <= 122 || x >= 48 && x <= 57 {
		if x >= 0 && x <= 47 || x >= 58 && x <= 64 || x >= 91 && x <= 96 || x > 122 {
			flag = false
		}
	}
	return flag
}

func IsNumeric(str string) bool {
	strA := []rune(str)
	flag := true

	for _, x := range strA {
		if x >= 0 && x <= 47 || x > 57 {
			flag = false
		}

	}
	return flag
}

func IsLower(str string) bool {
	strA := []rune(str)
	flag := true

	for _, x := range strA {
		if x >= 0 && x <= 96 || x >= 123 {
			flag = false
		}

	}
	return flag
}
