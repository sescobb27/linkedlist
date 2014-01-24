package main

import (
	"fmt"
	"linkedlist"
)

func main() {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append("Hola")
	fmt.Println(list.Head.Value)
	fmt.Println(list.Head.NextNode.Value)

	for node := list.Head; node != nil; node = node.NextNode {
		fmt.Println(node.Value)
	}
}
