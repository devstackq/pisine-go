package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	arg := os.Args[1:]
	g := 0 //
	for range arg {
		g++
	}

	//count args

	if g > 1 { // condition
		fmt.Println("Too many arguments")
	} else if g < 1 {
		fmt.Println("File name missing")
	} else {
		// content readfile from - args 1 -> filename.txt
		content, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			//if notfound file -> error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%s\n", content)
		}
	}
}
