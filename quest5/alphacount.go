package piscine

func AlphaCount(str string) int {

	arrayRune := []rune(str) // breaking up string - rune
	count := 0

	for _, char := range arrayRune { // extract each char, for id:0, from to[range]

		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' { // if char, - match - from 65 to 90 || 97 to 122
			count = count + 1

		}

	}
	return count
}
