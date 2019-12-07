package piscine

func BasicJoin(strs []string) string {
	str := ""

	for _, v := range strs {
		str = str + v
	}
	return str
}
