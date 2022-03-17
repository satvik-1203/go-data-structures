package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (self *LinkedList) addFirst(value int) {

	if self.Head == nil {
		self.Head = &Node{value, nil}
		self.Tail = self.Head
		return
	}

	node := Node{value, nil}
	node.next = self.Head
	self.Head = &node
}

func (self *LinkedList) addLast(value int) {

	if self.Head == nil {
		self.Head = &Node{value, nil}
		self.Tail = self.Head
		return
	}

	node := Node{value, nil}
	self.Tail.next = &node
	self.Tail = &node
}

func (self LinkedList) print() {
	current := self.Head
	fmt.Println()
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
}

func (self *LinkedList) removeFirst() {
	if self.Head == nil {
		return
	}

	self.Head = self.Head.next
}

func (self *LinkedList) removeLast() {

	if self.Head == nil {
		return
	}

	current := self.Head

	for current.next != self.Tail {

		current = current.next
	}

	current.next = nil
	self.Tail = current

}
