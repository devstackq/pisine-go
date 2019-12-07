package piscine

func FindNextPrime(nb int) int {

	if nb <= 1 {
		return 2
	}
	for i := 2; i <= nb/2; i++ { // [prime nums start with 2]
		if nb%i == 0 { // if 8 mod 2 ==0, -> recursion search next prime nums, 11 /5 != 0, return prime num
			return FindNextPrime(nb + 1) // 8+1, 9+1, 10+1
		}
	}
	return nb
}

//from recursion,
//find is prime numbers
