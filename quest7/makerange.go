package piscine

func MakeRange(min, max int) []int {

	var arrInt []int

	if min >= max {
		return arrInt
	}
	arrInt = make([]int, max-min) // size, leng array 5,

	for i := min; i < max; i++ { // 5 < 10,
		arrInt[i-min] = i // 5 - 5, 5 -4, 5-2
	}
	return arrInt
}
