package piscine

func Sqrt(nb int) int {

	var b int
	var c int

	for a := 1; a <= nb; a++ { // 25, 5 loop
		if a > 0 {
			b = nb / a               // b = 25/5 =5
			if nb%a == 0 && b == a { // 25/5 == 0 T, 5 == 5 T
				return b // 5
			}
		}
	}
	return c
}

// func Sqrt(nb int) int {
// 	if nb < 0 {
// 		return -1
// 	}
// 	if nb == 0 || nb == 1 {
// 		return nb
// 	}

// 	lowerbound := 1  // 1
// 	upperbound := nb // 16 // 8

// 	root := lowerbound + (upperbound-lowerbound)/2 //nb=16 //8, root 1.5

// 	for root > nb/root || root+1 <= nb/(root+1) { //8 ? 2 true | 1.5 ? 2 false 8> 16/ 2, 8 > 2 true
// 		if root > nb/root { //8 true
// 			upperbound = root // upper = 8
// 		} else {
// 			lowerbound = root
// 		}
// 		root = lowerbound + (upperbound-lowerbound)/2 // 1 +
// 	}
// 	return root
// }

// func Sqrt(nb int) int {
// 	if nb == 1 {
// 		return 1
// 	} else {
// 		for i := 1; i <= nb; i++ {
// 			if i*i == nb {
// 				return i
// 			}
// 		}

// 	}
// 	return 0
// }
