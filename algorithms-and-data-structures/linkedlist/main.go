package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (ll *LinkedList) Insert(data int) {
	newNode := Node{data: data, next: nil}
	if ll.head == nil {
		ll.length++
		ll.head = &newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (ll *LinkedList) RemoveHead() {
	if ll.head == nil {
		fmt.Println("Cannot remove empty linked list!")
	} else {
		ll.head = ll.head.next
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.Insert(1)
	linkedList.Insert(2)
	linkedList.RemoveHead()
	linkedList.RemoveHead()
	linkedList.RemoveHead()
	linkedList.Display()
}
