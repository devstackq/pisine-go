package main

import (
	"fmt"

	piscine ".."
)

func main() {
	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "2")
	piscine.ListPushFront(link, "you")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "mdsan")

	fmt.Println(piscine.ListSize(link))
}
