package main

import (
	"fmt"
)



type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushBack adds a new node with `data` at the end of the list `l`.
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Head == nil {
		// If the list is empty, Head and Tail both point to the new node
		l.Head = newNode
		l.Tail = newNode
	} else {
		// If list is not empty, append at Tail and move Tail pointer
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}


func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
