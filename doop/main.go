package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	arguments := os.Args[1:]

	var count int // Счетчик элементов //

	var n1, n2 string // Числа(string) //

	var action string // Операнд //

	var res, a, b int // результат 1, 2 аргумент  //

	var err string // Error //

	var strArr string

	var overflow int

	var zero int

	maxVal := 9223372036854775806
	minVal := -9223372036854775807

	for range arguments {
		count++
	}

	if count == 3 {
		if IsNbr(arguments) {

			for _, s := range arguments[0] { /* Число в строку*/
				n1 = n1 + string(s)
			}

			for _, s := range arguments[2] {
				n2 = n2 + string(s)
			}

			action = arguments[1]

			a = Atoi(n1)
			b = Atoi(n2)

			if ((n1 > "9223372036854775806" || n1 < "-9223372036854775807") && (a >= maxVal || a <= minVal)) || ((n2 > "9223372036854775806" || n2 < "-9223372036854775807") && (b >= maxVal || b <= minVal)) {
				overflow = 1

			}
			if (a > 0 && maxVal-a < 0) || (a < 0 && minVal-a > 0) || (b > 0 && maxVal-b < 0) || (b < 0 && minVal-b > 0) {
				overflow = 1

			}

			if action == "+" {

				res = a + b
			} else if action == "-" {

				res = a - b
			} else if action == "/" {
				if b == 0 {
					zero = 1
				} else {

					res = a / b
				}

			} else if action == "*" {

				res = a * b

			} else if action == "%" {
				if b == 0 {
					zero = 2
				} else {

					res = a % b
				}
			}

			if zero == 1 {
				err = "No division by 0"

				for _, v := range err {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
			} else if zero == 2 {
				err = "No modulo by 0"
				for _, v := range err {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')

			} else if overflow == 1 {
				z01.PrintRune('0')
				z01.PrintRune('\n')
			} else {

				strArr = Itoa(res)

				for _, v := range strArr {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')

			}

		} else if !(IsNbr(arguments)) {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		}

	} else {
		return
	}

}

func Atoi(str string) int {
	var num int
	var neg int = 1
	for i, v := range str {
		if i == 0 && v == '-' {
			neg = neg * -1

		} else {
			num = num * 10
			num = int(v-48) + num
		}
	}
	return num * neg
}

func Itoa(res int) string {
	var n int
	var str string
	var neg bool

	if res < 0 {
		res = res * -1
		neg = true
	}

	for i := 1; i > 0; i++ {
		if res < 10 {
			str = string(res+48) + str
			break
		}
		if res > 9 {
			n = res % 10
			res = res / 10
			str = string(n+48) + str
		}
	}
	if neg == true {
		return "-" + str
	}
	return str

}

func IsNbr(arg []string) bool {

	for _, v := range arg[0] {
		if (v < 48 || v > 57) && v != '-' {
			return false
		}
	}

	for _, v := range arg[2] {
		if (v < 48 || v > 57) && v != '-' {
			return false
		}
	}
	for _, v := range arg[1] {
		if v != '+' && v != '-' && v != '*' && v != '/' && v != '%' {
			return false
		}
	}
	return true

}
