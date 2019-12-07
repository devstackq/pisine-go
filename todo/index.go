package piscine

func Index(s string, toFind string) int {

	//counter 1
	var lenT, lenF int

	//count1
	for range s {
		lenT++
	}
	for range toFind {
		lenF++
	}
	if lenF == 0 || s == toFind {
		return 0 // condition by task
	}
	//toFine =3,
	//s = 10
	//match s == toFine
	//return index
	for i := 0; i < lenT-lenF; i++ { // compare lenText - lenFind, condition for - loop
		if s[i:i+lenF] == toFind { // match, lenFind, i[0] from end lenText
			return i
		}

	}
	return -1
}
