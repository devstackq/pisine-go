package piscine

//Write a function IsSorted that returns true if the slice of int is sorted, and that returns false otherwise.
func IsSorted(f func(a, b int) int, tab []int) bool {

	len1 := 0 // length - for 1 case, when a > b, = 1
	len2 := 0 // length - for 2 case, when a == b, = 0
	len3 := 0 //length - for 3 case, when a < b, = -1

	count := 0 // count all elemetns arrray int

	for range tab {
		count++
	} //6

	var flag bool // flag for return bool value, default false
	var res int   // save result - each call func f()

	for i := 0; i < count-1; i++ { // for loop, do it, before count -1
		res = f(tab[i], tab[i+1]) // start func f, 1 iterartion, func f(0,1), 2 iter f(1,2), ...
		if res == 0 {             // if func f, return 1, len + 1...
			len1++
		} else if res > 0 {
			len2++
		} else if res < 0 {
			len3++ //
		}
	}
	// comapare  res length,  with length count -1, if match return true, else return false
	if len1 == count-1 { //count = 6
		flag = true
	} else if len2 == count-1 {
		flag = true
	} else if len3 == count-1 {
		flag = true
	} else {
		flag = false
	}
	return flag
}

func f(a, b int) int { //{0, 1, 2, 3, 4, 5} -> -1,-1,-1,-1,-1, { 5,4,3,2,1,0} -> 1,1,1,1,1,
	// {4,7,2} -1,
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return -1 // -1, -1, -1, -1, -1
}
