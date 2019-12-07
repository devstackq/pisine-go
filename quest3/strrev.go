package piscine

func StrRev(s string) string {
	old := []rune(s)
	new := []rune(s)

	count := 0 // count leng //5

	for range old {
		count++
	}

	for i := 0; i < count; i++ { // execute, before i < count 5
		new[count-i-1] = old[i] //   new[5-0-1] = old[0], new 4 index = 0index old// new[5-1-1], 3idx = 1idx old
	}
	return string(new)
}
