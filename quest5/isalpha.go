package piscine

func IsAlpha(str string) bool {
	sym := []rune(str)

	flag := true

	for x := range sym {
		//a-z, A-Z, 0-9
		if sym[x] >= 0 && sym[x] <= 47 || sym[x] >= 58 && sym[x] <= 64 || sym[x] >= 91 && sym[x] <= 96 || sym[x] >= 123 {
			flag = false
			break
		}
	}
	return flag
}
