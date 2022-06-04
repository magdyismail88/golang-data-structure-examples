package main

import (
	"fmt"
	"errors"
)

type Node struct {
	Val int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Insert(value int) {
	if ll.Head == nil {
		ll.Head = &Node{Val: value,}
		return
	}

	ptr := ll.Head

	for ptr.Next != nil {
		ptr = ptr.Next
	}

	ptr.Next = &Node{Val: value}
	return
}

func (ll *LinkedList) InsertAt(index, value int) error {
	if index == 0 {
		if ll.Head == nil {
			ll.Head = &Node{Val: value,}
		} else {
			ll.Head = ll.Head.Next
			ll.Head = &Node{Val: value,}
		}
	}

	counter := 0
	ptr, currPtr := ll.Head, ll.Head

	for  ptr.Next != nil && counter != index {
		currPtr = ptr
		ptr = ptr.Next
		counter += 1
	}

	if ptr == nil {
		return errors.New("Index not found")
	}

	newNode := &Node{Val: value,}
	newNode.Next = ptr
	currPtr.Next = newNode

	return nil
}

func (ll *LinkedList) Remove(value int) error {

	if ll.Head == nil {
		return errors.New("Linked list is empty")
	}

	if value == ll.Head.Val {
		if ll.Head.Next == nil {
			ll.Head = nil
		} else {
			next := ll.Head.Next
			ll.Head = next
			ll.Head.Next = next.Next
		}
	}

	currPtr := ll.Head
	ptr := ll.Head

	for ptr.Next != nil && ptr.Val != value {
		currPtr = ptr
		ptr = ptr.Next
	}

	if ptr == nil {
		return errors.New("Input value not found")
	}

	if ptr.Next == nil {
		currPtr.Next = nil
		// ptr = nil
	} else {
		currPtr.Next = ptr.Next
		ptr = ptr.Next
	}
	return nil
}

func (ll *LinkedList) RemoveAt(index int) error {
	if index == 0 {
		if ll.Head == nil {
			return errors.New("The list is already empty")
		} else {
			ll.Head = ll.Head.Next
		}
	}

	ptr, currPtr := ll.Head, ll.Head
	counter := 0

	for ptr.Next != nil && counter != index {
		currPtr = ptr
		ptr = ptr.Next
		counter += 1
	}

	if ptr == nil {
		return errors.New("Invalid delete operation")
	}

	if ptr.Next == nil {
		currPtr.Next = nil
	} else {
		currPtr.Next = ptr.Next
		ptr = nil
	}

	return nil
}

func (ll *LinkedList) Values() []int {
	if ll.Head == nil {
		return nil
	}

	values := make([]int, 0)

	ptr := ll.Head

	values = append(values, ptr.Val)

	for ptr.Next != nil {
		ptr = ptr.Next
		values = append(values, ptr.Val)
	}

	return values
}

func main() {

	ll := new(LinkedList)

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(5)

	fmt.Println("Before insert...")
	for _, val := range ll.Values() {
		fmt.Println(val)
	}

	err := ll.InsertAt(3, 4)
	if err != nil {
		panic(err)
	}

	fmt.Println("After insert...")
	for _, val := range ll.Values() {
		fmt.Println(val)
	}

	err = ll.RemoveAt(3)
	if err != nil {
		panic(err)
	}

	fmt.Println("After deleted at...")
	for _, val := range ll.Values() {
		fmt.Println(val)
	}

	// err := ll.Remove(4)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, val := range ll.Values() {
	// 	fmt.Println(val)
	// }

}