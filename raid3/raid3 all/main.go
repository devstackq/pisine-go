package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arr []rune

	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		arr = append(arr, char)
	}
	str := string(arr)

	x := 0
	y := 0
	for _, char := range arr {
		if char != '\n' && y == 0 {
			x++
		}
		if char == '\n' {
			y++
		}
	}
	if x == 0 || y == 0 {
		fmt.Println("Not a Raid function")
		return
	}
	if checkF(str, x, y, 'o', 'o', 'o', 'o', '-', '|') {
		fmt.Printf("[raid1a] [%v] [%v]\n", x, y)
		return
	}

	if checkF(str, x, y, '/', '\\', '\\', '/', '*', '*') {
		fmt.Printf("[raid1b] [%v] [%v]\n", x, y)
		return
	}

	n := 0
	if checkF(str, x, y, 'A', 'A', 'C', 'C', 'B', 'B') {
		n++
		fmt.Printf("[raid1c] [%v] [%v]", x, y)
	}
	if checkF(str, x, y, 'A', 'C', 'A', 'C', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[raid1d] [%v] [%v]", x, y)
	}
	if checkF(str, x, y, 'A', 'C', 'C', 'A', 'B', 'B') {
		if n > 0 {
			fmt.Print(" || ")
		}
		n++
		fmt.Printf("[raid1e] [%v] [%v]", x, y)
	}

	if n > 0 {
		fmt.Println()
		return
	}

	fmt.Println("Not a Raid function")
}

func checkF(str string, x, y int, c1, c2, c3, c4, hor, ver rune) bool {
	var arrE []rune
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 {
				if j == 0 {
					arrE = append(arrE, c1)
				} else if j == x-1 {
					arrE = append(arrE, c2)
				} else {
					arrE = append(arrE, hor)
				}
			} else if i == y-1 {
				if j == 0 {
					arrE = append(arrE, c3)
				} else if j == x-1 {
					arrE = append(arrE, c4)
				} else {
					arrE = append(arrE, hor)
				}
			} else {
				if j == 0 || j == x-1 {
					arrE = append(arrE, ver)
				} else {
					arrE = append(arrE, ' ')
				}
			}
		}
		arrE = append(arrE, '\n')
	}
	strE := string(arrE)
	if strE == str {
		return true
	}
	return false
}
