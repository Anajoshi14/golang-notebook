package linkedlist

import "fmt"

type Node struct {
	num  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddToFront(n int) {
	if l.head == nil {
		l.head = &Node{num: n, next: nil}
	} else {
		newNode := &Node{num: n, next: l.head}
		l.head = newNode
	}

}

func (l *LinkedList) AddToBack(n int) {
	if l.head == nil {
		l.head = &Node{num: n, next: nil}
	} else {
		newNode := &Node{num: n, next: nil}
		temp := l.head
		for temp.next != nil {
			temp = temp.next //moving the temp pointer to the next element
		}
		temp.next = newNode

	}
}

func (ll *LinkedList) DisplayLinkedList() {

	temp := *ll.head

	for temp.next != nil {
		fmt.Println("->", temp.num)
		temp = *temp.next

	}

	fmt.Println("->", temp.num)

}
