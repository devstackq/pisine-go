package piscine

func Capitalize(s string) string {

	n := []rune(s)
	for i, value := range s {
		if n[i] >= 97 && n[i] <= 122 && i != 0 { // chek big 1 Capital symbol

			if n[i-1] < 48 || n[i-1] > 57 && n[i-1] < 65 || n[i-1] > 90 && n[i-1] < 97 || n[i-1] > 122 {
				n[i] = value - 32
			} //to upper, if left or right side, after n[i], - upper Letter,

		} else if n[i] >= 65 && n[i] <= 90 && i != 0 {
			if n[i-1] >= 48 && n[i-1] <= 57 || n[i-1] >= 97 && n[i-1] <= 122 || n[i-1] >= 65 && n[i-1] <= 90 {
				n[i] = value + 32
			} // to lower, if n[i], A, then lower letter
		} else if i == 0 {
			if value >= 97 && value <= 122 {
				n[i] = value - 32
			}
		}
	}
	return string(n)
}
