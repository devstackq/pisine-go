package piscine

func IsPrime(nb int) bool {
	var flag bool
	flag = true

	if nb <= 0 || nb == 1 {
		return false
	}

	for k := 2; k*k <= nb; k++ {
		if nb%k == 0 {
			flag = false
		}
	}
	return flag
}
