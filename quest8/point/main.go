package main

import (
	"github.com/01-edu/z01"
)

//1 decalare point, and type struct, same - class
type point struct {
	x string
	y string
}

//3 func call, init - ptr  take, x
// class - init variable - value
func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

//print points - as rune
func printPoints(arg string) {
	for _, v := range arg {
		z01.PrintRune(v)
	}
}

func main() {
	//2 assign variable - points, by -> pointer
	points := &point{}

	//3 set point
	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printPoints(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printPoints(points.y)
	z01.PrintRune('\n')
}
