package piscine

func IsUpper(str string) bool {
	sym := []rune(str)

	flag := true

	for x := range sym {
		if sym[x] >= 0 && sym[x] <= 66 || sym[x] >= 91 {
			flag = false
		}
	}
	return flag
}
