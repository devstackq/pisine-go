package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	count := 0
	for range os.Args {
		count++
	}
	//if not file ->
	for i := 1; i < count; i++ { //2,
		content, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			for _, v := range os.Args[i] { //if error, not file// show me hello world
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else {
			str := string(content)
			for _, v := range str { // give me content - inside file
				z01.PrintRune(v)
			}
			if i < count {
				z01.PrintRune('\n')
			}
		}
	}
}

// ./cat Hello There
// Hello
// There

//2 case leng =1
// else {
// 	/cat quest8.txt
// 	return
// 	Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing

// }

// 3 case {
// 	student@ubuntu:~/piscine-go/cat$ ./cat quest8.txt quest8T.txt
// 	return all contain -> 2 files
// }
