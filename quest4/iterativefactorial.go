package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 || nb > 20 {
		return 0
	}

	num := 1
	for i := 1; i <= nb; i++ {
		num = num * i
	}
	return num
}
