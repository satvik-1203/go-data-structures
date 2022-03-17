package main

import "fmt"

func main() {

	tree := Tree{}
	tree.addNode(10)
	tree.addNode(11)
	tree.addNode(9)
	fmt.Println(tree.height())
	tree.inOrder()
}
