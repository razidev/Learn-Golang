package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Aziz")
	data.PushBack("Syahputro")
	data.PushFront("Razi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
