package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Luthfi") // cara nambahin datanya gini
	data.PushBack("Izzuddin")
	data.PushBack("Hanif")

	first := data.Front()
	fmt.Println(first.Value)

	second := data.Front().Next()
	fmt.Println(second.Value)

	third := data.Front().Next().Next()
	fmt.Println(third.Value)

	// ambil data pake for
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
