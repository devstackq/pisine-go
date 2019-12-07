package piscine

func CollatzCountdown(start int) int {

	result := start //12

	if result <= 0 {
		return -1
	}

	if result <= 1 { // base condition
		return 1
	}
	// count 10
	// 12 % 2 ==0, 6, 3, 10, 5, 16-8, 4/2= 2, 1
	if result%2 == 0 { // even
		return 1 + CollatzCountdown(result/2)
	} // odd
	//10, 16
	return 1 + CollatzCountdown(3*result+1)

}
