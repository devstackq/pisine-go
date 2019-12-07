package piscine

import "fmt"


//2call,
func Any(f func(string) bool, arr []string) bool {

	var g bool
	// for range, take arr value, each word -> Hello -> and send isNumeric function, at the end loop
	
	for _, v := range arr { // 1 - Hello, 2 How, 3 are ..
		for 
		g = f(v) // start isNumeric func - word Hello
		if g == true {
			break
		}
	}

	return g // fasle
}

//3call - return res -> Any
// take string - Hello, and compare each letter -> and return result
// func IsNumeric(str string) bool {
// 	sym := []rune(str)

// 	for x := range sym {
// 		if sym[x] >= 0 && sym[x] <= 47 || sym[x] >= 58 {
// 			return false
// 		}
// 	}
// 	return true
// }
