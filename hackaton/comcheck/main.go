package main

import (
	"fmt"
	"os"
)

func main() {

	arg := os.Args

	for _, v := range arg {
		if v == "galaxy 01" || v == "galaxy" || v == "01" {
			fmt.Println("Alert!!!")
			break

		}
	}

}
