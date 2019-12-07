package piscine

func IsLower(str string) bool {
	sym := []rune(str)

	flag := true

	for x := range sym {
		if sym[x] >= 0 && sym[x] <= 96 || sym[x] >= 123 {
			flag = false
		}
	}
	return flag
}
