package piscine

func SortIntegerTable(table []int) {

	leng := len(table) // index from start 0

	for i := 1; i < leng; i++ { // 0
		for j := 0; j < leng-i; j++ { // j0
			if table[j] > table[j+1] {
				tmp := table[j]
				table[j] = table[j+1]
				table[j+1] = tmp
			}
		}
	}
}
