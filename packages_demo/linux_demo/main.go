package main

import (
	"container/list"
	"fmt"
)

func main() {

	listStru := list.List{}

	root := list.Element{Value: "root"}
	root2 := list.Element{Value: "root2"}
	root3 := list.Element{Value: "root3"}

	listStru.PushFront(root)
	listStru.PushBack(root2)
	listStru.PushBack(root3)

	//listStru.InsertAfter(root, listStru.Front())

	fmt.Println(listStru)

	fmt.Println(listStru.Back().Next())
	fmt.Println(listStru.Back().Prev())

	for node := listStru.Front(); node != nil; node = node.Next() {
		fmt.Println("Node", node.Value)
	}
}
