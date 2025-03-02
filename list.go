package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Syafiq")
	data.PushBack("Fajrian")
	data.PushBack("Emha")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Syafiq

	next := head.Next() // Fajrian
	fmt.Println(next.Value)

	next = next.Next() // Emha
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
