package piscine

func ToUpper(s string) string {

	byteArr := []byte(s) //arrayt byte [0:44, 1: 66 ...]

	for idx := range byteArr { // index byte
		if byteArr[idx] >= 97 && byteArr[idx] <= 122 { //compare 0:H-72, 1:e-101 ...
			byteArr[idx] = byteArr[idx] - 32 // current index value, minus - 32, , lower - to upper 101 -32,
		}
	}
	return string(byteArr) // return all data, array
}
