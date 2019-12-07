package piscine

func Fibonacci(index int) int {

	var result int

	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {

		result = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return result
}
