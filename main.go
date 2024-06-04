package main

import (
	"fmt"

	"github.com/ordili/gods/lists/singlelinkedlist"
)

func main() {
	list := singlelinkedlist.New[int](1, 2, 3)
	fmt.Println(list)
	list.Remove(0)
	fmt.Println(list)
	list.Remove(1)
	fmt.Println(list)
	list.Remove(0)
	fmt.Println(list)
	fmt.Println("Done")
}
