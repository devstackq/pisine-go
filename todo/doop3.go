package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if stringArrLen(args) != 3 { //1
		return
	}

	firstArg := atoi(args[0])  //2
	secondArg := atoi(args[2]) //2
	operator := args[1]

	switch operator { //3
	case "*": //3
		fmt.Println(firstArg * secondArg) //4
	case "/": //3
		if secondArg == 0 {
			fmt.Println("No division by 0")
			return
		}

		fmt.Println(firstArg / secondArg) //4
	case "%": //3
		if secondArg == 0 {
			fmt.Println("No Modulo by 0")
			return
		}

		fmt.Println(firstArg % secondArg) //4
	case "+": //3
		fmt.Println(firstArg + secondArg) //4
	case "-": //3
		fmt.Println(firstArg - secondArg) //4
	default:
		fmt.Println(0)
	}
}

func stringArrLen(arr []string) int {
	count := 0
	for range arr {
		count++
	}

	return count
}

func atoi(s string) int {
	if strLen(s) == 0 {
		return 0
	}

	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if strLen(s) < 1 {
			return 0
		}
	}

	nm := 0

	for _, ch := range s {
		if !containsIn0to9(ch) {
			fmt.Println(0)
			os.Exit(0)
		}
		nm = nm*10 + charToInt(ch)
	}

	if s0[0] == '-' {
		nm *= -1
	}
	return nm
}

func containsIn0to9(ch rune) bool {
	for i := '0'; i <= '9'; i++ {
		if ch == i {
			return true
		}
	}

	return false
}

func charToInt(ch rune) int {
	count := 0
	if ch < 48 && ch > 58 {
		return 0
	}

	for i := '1'; i <= ch; i++ {
		count++
	}

	return count
}

func strLen(s string) int {
	count := 0
	for range s {
		count++
	}

	return count
}
