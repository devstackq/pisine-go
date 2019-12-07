package piscine

func AppendRange(min, max int) []int {

	var arrInt []int

	if min == max || max < min {
		return arrInt
	}

	// array int
	if max > min { // 10 > 5, else return 0
		for i := min; i < max; i++ { // 5 < 5 10-5, 5,
			arrInt = append(arrInt, i) //[5,6,7,8,9 ]
		}
	}
	return arrInt
}

//1 array empy int

