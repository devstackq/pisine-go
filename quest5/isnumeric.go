package main

func IsNumeric(str string) bool {
	sym := []rune(str)

	flag := true
	for x := range sym {
		if sym[x] >= 0 && sym[x] <= 47 || sym[x] >= 58 {
			flag = false
		}
	}
	return flag
}
