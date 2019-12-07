package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	str := os.Args
	t := 0
	a := 0
	b := 0
	c := ""

	var boo bool
	var boo1 bool

	for range str {
		t++
	}

	if t > 4 || t < 4 {
		fmt.Println(0)
	}

	if t == 4 {
		// if str[1] > "9223372036854775807" || str[3] > "9223372036854775807" || str[1] < "-9223372036854775808" || str[3] < "9223372036854775808" {
		//     fmt.Println(0)
		// } else if str[1] < "9223372036854775807" && str[1] < "-9223372036854775808" || str[1] < "-9223372036854775808" || str[3] < "9223372036854775808"

		a, boo1 = Convert(str[1])
		b, boo = Convert(str[3])
		fmt.Println(a)
		c = Operation(str[2], a, b, boo1, boo)

		for _, k := range c {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}
}

func Convert(st string) (int, bool) {

	letterE := false
	str := []rune(st)
	n := 0
	if str[0] != '-' {
		for k := range str {
			if str[k] < '0' || str[k] > '9' {
				letterE = true
				return 0, false
			}
			if (str[k] >= '0' && str[k] <= '9') && !letterE {
				n = n*10 + int(str[k]-48)
			}
		}
		if letterE == false {
			return n, true
		}
	}

	if str[0] == '-' {
		for _, k := range str[1:] {
			if k < '0' || k > '9' {
				letterE = true
				return 0, false
			}
			if (k >= '0' && k <= '9') && !letterE {
				n = n*10 + int(k-48)
			}
		}
		if letterE == false {
			return n * -1, true
		}
	}
	return 0, false
}

func Operation(str string, a, b int, boolean1, boolean2 bool) string {
	t := 0
	c := 0
	sl := ""
	for range str {
		t++
	}

	if t == 1 {
		//add
		if str == "+" && boolean1 && boolean2 {
			c = a + b
			if c > 0 {
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return sl
			} else if c < 0 {
				c = c * -1
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return "-" + sl
			} else if c == 0 {
				return "0"
			}
		}
		//minus
		if str == "-" && boolean1 && boolean2 {
			c = a - b
			if c > 0 {
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return sl
			} else if c < 0 {
				c = c * -1
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return "-" + sl
			} else if c == 0 {
				return "0"
			}
		}
		//dev
		if str == "/" && b != 0 && boolean1 && boolean2 {
			c = a / b
			if c > 0 {
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return sl
			} else if c < 0 {
				c = c * -1
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return "-" + sl
			} else if c == 0 {
				return "0"
			}
		}
		if str == "/" && b == 0 && boolean1 && boolean2 {
			return "No devision by 0"
			//mod
		} else if str == "%" && b != 0 && boolean1 && boolean2 {
			c = a % b
			if c > 0 {
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return sl
			} else if c < 0 {
				c = c * -1
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return "-" + sl
			} else if c == 0 {
				return "0"
			}
		}
		if str == "%" && b == 0 && boolean1 && boolean2 {
			return "No modulo by 0"
		}
		//multiple
		if str == "*" && boolean1 && boolean2 {
			c = a * b
			if c > 0 {
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return sl
			} else if c < 0 {
				c = c * -1
				for c > 0 {
					sl = string(c%10+48) + sl
					c = c / 10
				}
				return "-" + sl
			} else if c == 0 {
				return "0"
			}
		} else {
			return "0"
		}
	}
	return "0"
}
