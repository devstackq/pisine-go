package main

import "fmt"

func main() {
	arrInt := []int{3, 234, 6, 2, 4, 7, 5, 0}
	var res int
	res = Max(arrInt)
	fmt.Print(res)
}

func Max(arr []int) int {

	//sort
	//print max value

	count := 0

	for range arr {
		count++
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[count-1]

}
