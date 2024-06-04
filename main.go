package main

import (
	"fmt"

	"github.com/ordili/gods/maps/hashmap"
)

func main() {
	m := hashmap.New[string, int]()
	m.Put("JD", 100)
	m.Put("QQ", 200)
	fmt.Println(m)
	fmt.Println("Done")
}
