package linkedlist

import (
	"errors"
)

type Item interface{}

type Node struct {
	Value    Item
	NextNode *Node
}

type LinkedList struct {
	Head *Node
	Node *Node
	Size int32
}

func (list *LinkedList) Prepend(item Item) {
	if list.Size == 0 {
		list.Node = &Node{Value: item}
		list.Head = list.Node
		list.Size++
	} else {
		tmp := list.Head
		list.Head = &Node{Value: item}
		list.Head.NextNode = tmp
		list.Size++
	}
}

func (list *LinkedList) Append(item Item) {
	if list.Size == 0 {
		list.Node = &Node{Value: item}
		list.Head = list.Node
		list.Size++
	} else {
		list.Node.NextNode = &Node{Value: item}
		list.Size++
		list.Node = list.Node.NextNode
	}
}

func (list *LinkedList) Each(f func(node *Node)) {
	for node := list.Head; node != nil; node = node.NextNode {
		f(node)
	}
}

func (list *LinkedList) Find(item Item) (Item, bool) {
	for node := list.Head; node != nil; node = node.NextNode {
		if node.Value == item {
			return node, true
		}
	}
	return nil, false
}

func (list *LinkedList) Remove(item Item) error {
	if list.Size == 0 {
		return errors.New("List is empty")
	}

	if list.Head.Value == item {
		list.Head = list.Head.NextNode
		list.Size--
		return nil
	}
	remove := false
	last_item := list.Head
	if_exist_remove := func(node *Node) {
		if node.Value == item {
			last_item.NextNode = node.NextNode
			remove = true
			list.Size--
		} else {
			last_item = node
		}
	}
	list.Each(if_exist_remove)
	var itemError error
	if !remove {
		itemError = errors.New("Item not found")
	}
	return itemError
}

func (list *LinkedList) Clear() {
	list.Size = 0
	list.Head = nil
	list.Node = nil
}
