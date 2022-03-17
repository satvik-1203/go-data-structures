package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int

	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (self *Tree) addNode(value int) {

	if self.root == nil {
		self.root = &Node{value: value}
		return
	}

	current := self.root

	for {
		if current.value > value && current.left == nil {
			current.left = &Node{value: value}
			break
		}

		if current.value <= value && current.right == nil {
			current.right = &Node{value: value}
			break
		}

		if current.value > value {
			current = current.left
		} else {
			current = current.right
		}

	}
}

func height(node *Node) int {

	if node == nil {
		return 0
	}

	return 1 + int(math.Max(float64(height(node.left)), float64(height(node.right))))
}

func (self Tree) height() int {
	return height(self.root)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.value)
	inOrder(node.right)
}

func (self Tree) inOrder() {
	inOrder(self.root)
}
