package piscine

func RecursivePower(nb int, power int) int {

	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	} else {
		return nb * RecursivePower(nb, power-1)
	}

}

// 4, 3
// 4 - 4*4*4
