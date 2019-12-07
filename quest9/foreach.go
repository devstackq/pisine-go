package piscine

// import "fmt"

// // start func forEach, (array int params) 1
// func main() {
// 	arr := []int{5, 2, 3, 3, 5, 6}
// 	ForEach(PrintNbr, arr)
// }

// 2 func call
func ForEach(f func(int), arr []int) {
	//arr = f(arr)
	for _, k := range arr { //recursion call func, take each value from - array int[] , and start func
		f(k)
	}

}

// // 3 func call
// func PrintNbr(n int) {
// 	sl := ""
// 	sl = sl + string(n+48) //5+48 = 53 -> "5"
// 	fmt.Print(sl)          // "1+2" "1+2+3+4+5+6" concat
// }

// // from - array int - take - just int value
