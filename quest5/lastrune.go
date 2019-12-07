package piscine

func LastRune(s string) rune {
	runeVal := []rune(s) // "S o m a" ["0: 33, 1:29, 2:67, 3:99"]
	count := 0

	for idx := range runeVal { // [0,1,2,3], 1,2,3,4
		count = idx // // 3+1 4,
	}
	return runeVal[count] // [4-1] runeVal[3] // a - last word
}

//return end val
