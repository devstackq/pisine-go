package piscine

func IterativePower(nb int, power int) int {

	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}

	valnb := nb                  // 4, 16, 64
	for i := 1; i < power; i++ { // 1/2/3
		nb = nb * valnb // 4 * 4, // 16 = 16 * 4
	}
	return nb //64
}
