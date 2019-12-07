package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {

	//alph := "abcdefghijklmnopqrstuvwxyz"
	var alph string

	// create own alphabet
	for i := 97; i < 122; i++ {
		alph = alph + string(i)
	}

	fmt.Println(alph)

	// create variable - type rune
	var nv rune

	for idx, v := range alph { // 0,1,2,.. 25
		if idx%2 == 1 { // 0 % 2== 1, //B 1 == 1, 3 % 2 = 1, 1 == 1
			nv = (v - 32)     // nv = b -> 98 - 32 = 66 -> B,  d - > 100 - 32 = 68 -> D
			z01.PrintRune(nv) // B, D
		} else {
			z01.PrintRune(v) // a , c
		}

	}
}
