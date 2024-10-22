package main

import (
	linkedlist "github.com/Anajoshi14/golang_repo/linked_list"
)

func main() {
	ll := linkedlist.LinkedList{}
	ll.AddToFront(7)
	ll.AddToFront(8)
	ll.AddToFront(9)
	ll.AddToBack(5)
	ll.AddToBack(6)
	ll.DisplayLinkedList()
}

//9 8 7 5 6
