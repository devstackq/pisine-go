package piscine

func Compact(ptr *[]string) int {

	lengArr := 0

	for _, v := range *ptr { // count not empty values, 1 step
		if v != "" {
			lengArr++ // 3
		}
	}
	// create -  array, with size
	compactArr := make([]string, lengArr) // 3 leng arr

	idx := 0
	for _, v := range *ptr { // add new in arr - each  value, from *ptr, not empty value
		if v != "" {
			compactArr[idx] = v
			idx++
		}

	}
	*ptr = compactArr // assign, value address
	return idx
}

// return  count elemetns - not nil value
// delete elemetns with slice
