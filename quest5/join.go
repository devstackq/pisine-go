package piscine

func Join(strs []string, sep string) string {

	str := ""

	for idx, val := range strs {
		if idx != 0 {
			str += sep
		}
		str = str + val
	}
	return str
}
