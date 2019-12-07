package piscine

func Rot14(str string) string {

	var text string

	for _, v := range str {
		if (v >= 'a' && v <= 'l') || (v >= 'A' && v <= 'L') {
			text = text + string(v+14)
		} else if (v >= 'm' && v <= 'z') || (v >= 'M' && v <= 'Z') {
			text = text + string(v-12)
		} else {
			text = text + string(v)
		}

	}
	return text
}
