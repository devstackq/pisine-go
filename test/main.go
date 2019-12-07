package main

import (
	"fmt"

	piscine ".."
)

func main() {
	arr := []int{1, 2, 2, 3, 3, 5, 5, 1, 4, 4, 4}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}
