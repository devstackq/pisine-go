package piscine

func StrLen(str string) int {
	text := []rune(str)
	var length int
	for i := range text {
		length = i
	}
	return length + 1
}
