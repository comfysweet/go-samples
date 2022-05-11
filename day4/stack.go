package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	tail *node
}

func main() {
	node1 := &node{data: 20}
	node2 := &node{data: 25}
	node3 := &node{data: 26}
	list := linkedList{}
	list.Push(node1)
	list.Push(node2)
	list.Push(node3)
	fmt.Println(list)
	list.Pop()
	fmt.Println(list)
	fmt.Println(list.Get().data)
	fmt.Println(list)
}

func (l linkedList) String() string {
	var str string
	for l.head != nil {
		str += fmt.Sprintf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	return str
}

func (l *linkedList) Push(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.next = l.head
		l.head = n
	}
}

func (l *linkedList) Pop() node {
	if l.head == nil {
		return node{}
	}
	n := *l.head
	l.head = l.head.next
	return n
}

func (l *linkedList) Get() node {
	if l.head == nil {
		return node{}
	}
	return *l.head
}
