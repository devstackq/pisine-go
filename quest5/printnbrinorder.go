package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {

	arrayNums := []int{} // [   ]
	count := 0
	if n == 0 {
		z01.PrintRune('0')
	}
	if n > 0 {
		for i := 0; i <= 100; i++ {
			if n > 0 {
				mod := n % 10                      //7
				arrayNums = append(arrayNums, mod) //  [] =  [7, ] // [7, ] = [7,2] , [7,2] = [..]
				n = n / 10
				count++
			}
		}
	}
	// fmt.Println( arrayNums)
	//nbr2
	for k := '0'; k <= '9'; k++ { // 48 -58
		for l := 0; l < count; l++ {
			if rune(arrayNums[l]+48) == k { // if match 'rune(0-9)' 3 == 3 && array[l] 3 + 48 = 51, 53 ==53
				//[3,   ] == [ 2 ] [50 == 2 = (48) ] ''
				z01.PrintRune(rune(arrayNums[l]) + 48)
			}
		}
	}
}
