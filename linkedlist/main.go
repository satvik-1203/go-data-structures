package main

func main() {
	linkedList := LinkedList{nil, nil}

	linkedList.addFirst(10)
	linkedList.addLast(20)
	linkedList.print()
	linkedList.removeLast()
	linkedList.print()
}
