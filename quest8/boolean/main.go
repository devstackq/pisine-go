package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	argVal := os.Args[1:]
	lengthOfArg := 0

	for range argVal { // 2
		lengthOfArg++
	}

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

// 4 func call
func printStr(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

//3 func call
func even(nbr int) int { //2
	if nbr%2 == 0 { // 2
		return 1
	} else {
		return 0
	}
}

//2func  call
func isEven(nbr int) bool { //2
	if even(nbr) == 1 {
		return true
	}
	return false
}

//1func call
