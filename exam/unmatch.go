package piscine

func Unmatch(arr []int) int {

	//count, 3
	// check
	//sort bubble
	// comapre - counter
	//if  match = -1
	//if i != -1 {
	//arr[i]
	l := 0
	for range arr {
		l++
	}
	// bubble sort
	//11,2,2,33,44, 4, 55, 4 ...
	for i := 1; i < l; i++ { // 0
		for j := 0; j < l-1; j++ { // j0
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	//11,,44,66,77, ->  -1 3

	// 1 =-1, 1 -1, -
	// -1,-1-1-1-1-1,  4
	for i := 0; i < l-1; i++ {
		if arr[i] == arr[i+1] {
			arr[i] = -1
			arr[i+1] = -1
		}
	}

	// if not -1, write value -> 4
	for i := 0; i < l; i++ {
		if arr[i] != -1 {
			return arr[i]
		}
	}
	return -1
}
