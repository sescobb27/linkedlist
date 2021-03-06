package linkedlist

import (
	"linkedlist"
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append("Hola")

	if list.Size != 2 {
		t.Error("Size should be 2")
	}

	if list.Head == nil {
		t.Error("List Head shouldn't be nil")
	}

	if list.Head.NextNode == nil {
		t.Error("List Header NextNode shouldn't be nil")
	}

	if list.Head == list.Head.NextNode {
		t.Error("List Head and List Head NextNode shouldn't be equal")
	}

	if list.Head.Value != 2 || list.Head.NextNode.Value != "Hola" || list.Head.NextNode != list.Node {
		t.Error("Values doesn't match")
	}
}

func TestMap(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append(5)
	list.Append(34)
	counter := 0
	to_s := func(node *linkedlist.Node) {
		// convert an interface into string
		node.Value = reflect.ValueOf(node.Value).String()
		if str_v, ok := node.Value.(string); ok {
			node.Value = str_v + " to_s functions"
			counter++
		} else {
			t.Error("can't convert to strings ", str_v, ok)
		}
	}
	list.Map(to_s)
	if counter != 3 {
		t.Error("counter should be 3")
	}
}

func TestPrepend(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append(5)
	list.Append(34)

	list.Prepend(100)
	if list.Head.Value != 100 {
		t.Error("List head should be 100")
	}

	if list.Head.NextNode.Value != 2 {
		t.Error("List head next node should be 2")
	}

	if list.Size != 4 {
		t.Error("List size should be 4")
	}

	if list.Head.NextNode == list.Node {
		t.Error("List Head shouldn't be the last element")
	}
}

func TestRemove(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(34)
	list.Prepend(100)
	var err error
	// remove head
	err = list.Remove(100)
	if err != nil {
		t.Error(err)
	} else {
		if list.Size != 7 {
			t.Error("List size should be 7 and is ", list.Size)
		}

		if _, ok := list.Find(100); ok {
			t.Error("100 shouldn't be in list")
		}
	}
	// remove tail
	err = list.Remove(34)
	if err != nil {
		t.Error(err)
	} else {
		if list.Size != 6 {
			t.Error("List size should be 6 and is ", list.Size)
		}

		if _, ok := list.Find(34); ok {
			t.Error("34 shouldn't be in list")
		}
	}
	// remove middle
	err = list.Remove(5)
	if err != nil {
		t.Error(err)
	} else {
		if list.Size != 5 {
			t.Error("List size should be 5 and is ", list.Size)
		}

		if _, ok := list.Find(5); ok {
			t.Error("5 shouldn't be in list")
		}
	}
}

func TestRemoveEmptyList(t *testing.T) {
	list := new(linkedlist.LinkedList)
	var item linkedlist.Item
	err := list.Remove(item)

	if err == nil {
		t.Error("Error should be 'List is empty'")
	}
}

func TestFind(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(34)
	list.Prepend(100)

	node, ok := list.Find(2)
	if !ok || node == nil {
		t.Error("Node shouldn't be nil and Ok should be true")
	} else if node.Value != 2 {
		t.Error("Error node value should be", 2)
	}

	node, ok = list.Find(5)
	if !ok || node == nil {
		t.Error("Node shouldn't be nil and Ok should be true")
	} else if node.Value != 5 {
		t.Error("Error node value should be", 5)
	}

	node, ok = list.Find(100)
	if !ok || node == nil {
		t.Error("Node shouldn't be nil and Ok should be true")
	} else if node.Value != 100 {
		t.Error("Error node value should be", 100)
	}

	node, ok = list.Find(500)
	if node != nil || ok {
		t.Error("Node should be nil and Ok should be false")
	}
}

func TestClear(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(34)
	list.Prepend(100)

	list.Clear()

	if list.Size != 0 || list.Head != nil || list.Node != nil {
		t.Error("List should be empty", list.Size, list.Head, list.Node)
	}
}

func TestGet(t *testing.T) {
	list := new(linkedlist.LinkedList)
	list.Append(2)
	list.Append("Hola")
	list.Append(4)
	list.Append("Golang")
	list.Append(6)
	list.Append(7)
	list.Append(34)
	list.Prepend("Last")

	// first item
	node, _err := list.Get(0)
	if node == nil || _err != nil {
		t.Error("Item not found", _err)
	} else if node.Value != "Last" {
		t.Error("Value should be", "Last")
	}
	// middle node
	node, _err = list.Get(1)
	if node == nil || _err != nil {
		t.Error("Item not found", _err)
	} else if node.Value != 2 {
		t.Error("Value should be", 2, node.Value)
	}
	node, _err = list.Get(4)
	if node == nil || _err != nil {
		t.Error("Item not found", _err)
	} else if node.Value != "Golang" {
		t.Error("Value should be", "Golang", node.Value)
	}
	// last node
	node, _err = list.Get(7)
	if node == nil || _err != nil {
		t.Error("Item not found", _err)
	} else if node.Value != 34 {
		t.Error("Value should be", 34)
	}
	// index out of bounds
	node, _err = list.Get(9)
	if node != nil || _err == nil {
		t.Error("Error should be Index Out of Bounds", _err)
	}
}

func TestConcat(t *testing.T) {
	numbers := new(linkedlist.LinkedList)
	numbers.Append(2)
	numbers.Append(3)
	numbers.Append(4)
	numbers.Append(5)
	numbers.Append(6)
	numbers.Append(7)
	numbers.Append(34)
	numbers.Prepend(100)

	words := new(linkedlist.LinkedList)
	words.Append("Hola")
	words.Append("Golang")
	words.Prepend("Last")

	numbers.Concat(words)

	if numbers.Size != 11 {
		t.Error("List Size should be the sum of both list")
	}

	if numbers.Head.Value != 100 {
		t.Error("List Head should be numbers head")
	}

	if numbers.Node.Value != "Golang" {
		t.Error("List Tails should be words tail")
	}

	words.Clear()

	if numbers.Size != 11 {
		t.Error("List Size should be the sum of both list")
	}

	if numbers.Head.Value != 100 {
		t.Error("List Head should be numbers head")
	}

	if numbers.Node.Value != "Golang" {
		t.Error("List Tails should be words tail")
	}

	if words.Size != 0 || words.Head != nil || words.Node != nil {
		t.Error("List should be empty", words.Size, words.Head, words.Node)
	}
}
