package piscine

func Abort(a, b, c, d, e int) int {

	table := []int{a, b, c, d, e} //5

	for i := 1; i < 5; i++ { // 0
		for j := 0; j < 4; j++ { // j0
			if table[j] > table[j+1] {
				tmp := table[j]
				table[j] = table[j+1]
				table[j+1] = tmp
			}
		}
	}
	return table[2]
}
