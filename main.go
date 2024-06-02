package main

import (
	"fmt"

	"github.com/ordili/gods/list/arraylist"
)

func main() {
	list := arraylist.New[int](1)
	fmt.Println(list)
	fmt.Println("Done")
}
