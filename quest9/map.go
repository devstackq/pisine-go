package piscine

// import (
// 	"fmt"
// )

// //1 call
// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6}
// 	result := Map(IsPrime, arr)
// 	fmt.Println(result)
// }

//2 call
func Map(f func(int) bool, arr []int) []bool {
	// define length 6
	len := 0
	for range arr {
		len++
	}
	var ms []bool = make([]bool, len) // ms[] - prepare variable - size, for write - answer isPrime func
	// call func isPrime and - write result, -> ms[0] ...etc
	for i, v := range arr {
		ms[i] = f(v)
	}
	return ms
	// return all result, ms[0 false, 1 true ...]

}

//2 call, take value -> Map func, and return result - bool
// func IsPrime(nb int) bool {
// 	var flag bool
// 	flag = true

// 	if nb <= 0 || nb == 1 {
// 		return false
// 	}

// 	for k := 2; k*k <= nb; k++ {
// 		if nb%k == 0 {
// 			flag = false
// 		}
// 	}
// 	return flag
// }
