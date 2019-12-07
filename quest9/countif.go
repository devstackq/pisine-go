package piscine

// count - number

func CountIf(f func(string) bool, tab []string) int {

	var g bool
	count := 0

	// for range, take arr value, each word -> Hello -> and send isNumeric function, at the end loop
	for _, v := range tab { // 1 - Hello, 2 How, 3 are ..
		g = f(v)       // start isNumeric func - word Hello
		if g == true { // count - true, and return  number, and return all nums
			count++
		}
	}
	return count
}

// func IsNumeric(str string) bool {
// 	sym := []rune(str)

// 	for x := range sym {
// 		if sym[x] >= 0 && sym[x] <= 47 || sym[x] >= 58 {
// 			return false
// 		}
// 	}
// 	return true
// }
