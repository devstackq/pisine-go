package piscine

func ActiveBits(n int) uint {

	if n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}

	var count uint = 1

	if n > 0 {
		if n%2 == 1 {
			return count + ActiveBits(n/2)

		} else if n%2 == 0 {
			return ActiveBits(n / 2)
		}

	}

	return count
}

// 	var count uint = 0
// 	//5
// 	var flag bool
// 	if n < 0 {
// 		flag = true
// 		n = n * -1
// 	}
// 	for n > 0 {
// 		if n%2 == 1 {
// 			count++
// 		}
// 		n = n / 2

// 	}
// 	if flag {
// 		count = count + 1
// 	}
// 	return count
// }
