package piscine

func Max(arr []int) int {

	l := len(arr)

	for i := 1; i < l; i++ { // 0
		for j := 0; j < l-1; j++ { // j0
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr[l-1]
}
