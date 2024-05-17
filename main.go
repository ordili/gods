package main

import "fmt"
import "github.com/ordili/gods/list/arraylist"

func main() {
	list := arraylist.New[int](1, 2, 3)
	list.Add(4, 5, 6)
	list.Remove(3)
	for {
		val, ret := list.Remove(0)
		if !ret {
			break
		}
		fmt.Println(val)
	}

	fmt.Println("Done")
}
