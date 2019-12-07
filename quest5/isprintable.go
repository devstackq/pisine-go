package piscine

func IsPrintable(str string) bool {
	sym := []rune(str)

	flag := true

	for x := range sym {
		if sym[x] >= 0 && sym[x] <= 31 {
			flag = false
		}
	}
	return flag
}
