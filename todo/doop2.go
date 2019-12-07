package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argOs := os.Args // v1-5 + v2-9

	pos := 0
	v1 := 0
	v2 := 0
	resCalc := ""

	var checkNum bool
	var checkNum2 bool

	for range argOs {
		pos++
	}

	if pos > 4 || pos < 4 {
		fmt.Println(0)
	}

	if pos == 4 {
		v1, checkNum = StrToInt(argOs[1])
		v2, checkNum2 = StrToInt(argOs[3])
		fmt.Println(v1)
		resCalc = Calculate(argOs[2], v1, v2, checkNum, checkNum2)

		for _, v := range resCalc {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

// string  nums convert -> to nums integer, else false
func StrToInt(arg string) (int, bool) {
	var letter bool //false
	runs := []rune(arg)
	num := 0
	// check letter, if letter, return 0, false
	if arg[0] != '-' {
		for i := range runs {
			if arg[i] < '0' || arg[i] > '9' {
				letter = true
				return 0, false
			} // convert str - to int
			if (arg[i] >= '0' && arg[i] <= '9') && !letter {
				num = num*10 + int(arg[i]-48)
			}

		}
		if letter == false {
			return num, true
		}
	}
	// minus convert - positive
	if arg[0] == '-' {
		for _, v := range arg[1:] {
			if v < '0' || v > '9' {
				letter = true
				return 0, false
			}
			if (v >= '0' && v <= '9') && !letter {
				num = num*10 + int(v-48)
			}
		}
		if letter == false {
			return num * -1, true
		}
	}
	return 0, false
}

func Calculate(operand string, v1, v2 int, checkNum, checkNum2 bool) string {

	lengArg := 0
	res := 0
	intToStr := ""

	for range operand {
		lengArg++
	}

	if lengArg == 1 {
		//add
		if operand == "+" && checkNum && checkNum2 {
			res = v1 + v2 // 5+7=12
			if res > 0 {
				for res > 0 {
					intToStr = string(res%10+48) + intToStr // 2+48 -> string("2")
					res = res / 10
				}
				return intToStr
				// negative value - positive
			} else if res < 0 {
				res = res * -1
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return "-" + intToStr

			} else if res == 0 {
				return "0"
			}
		}
		// minus

		if operand == "-" && checkNum && checkNum2 {
			res = v1 - v2 // 5+7=12
			if res > 0 {
				for res > 0 {
					intToStr = string(res%10+48) + intToStr // 2+48 -> string("2")
					res = res / 10
				}
				return intToStr
				// negative value - positive
			} else if res < 0 {
				res = res * -1
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return "-" + intToStr

			} else if res == 0 {
				return "0"
			}
		}

		if operand == "/" && v2 == 0 && checkNum && checkNum2 {
			return "No division by 0"

		} else if operand == "/" && checkNum && checkNum2 {
			res = v1 / v2 // 5+7=12
			if res > 0 {
				for res > 0 {
					intToStr = string(res%10+48) + intToStr // 2+48 -> string("2")
					res = res / 10
				}
				return intToStr
				// negative value - positive
			} else if res < 0 {
				res = res * -1
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return "-" + intToStr

			} else if res == 0 {
				return "0"
			}
		}

		//mod
		if operand == "%" && v2 != 0 && checkNum && checkNum2 {
			res = v1 % v2
			if res > 0 {
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return intToStr
			} else if res < 0 {
				res = res * -1
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return "-" + intToStr

			} else if res == 0 {
				return "0"

			}
		}
		if operand == "%" && v2 == 0 && checkNum && checkNum2 {
			return "No modulo by 0"
		}
		//multiply
		if operand == "*" && checkNum && checkNum2 {
			res = v1 * v2 // 5+7=12
			if res > 0 {
				for res > 0 {
					intToStr = string(res%10+48) + intToStr // 2+48 -> string("2")
					res = res / 10
				}
				return intToStr
				// negative value - positive
			} else if res < 0 {
				res = res * -1
				for res > 0 {
					intToStr = string(res%10+48) + intToStr
					res = res / 10
				}
				return "-" + intToStr

			} else if res == 0 {
				return "0"
			}
		} else {
			return "0"

		}

	}
	return "0"
}
