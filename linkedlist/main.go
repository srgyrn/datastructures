package main

import (
	"fmt"
	"github.com/srgyrn/datastructures/linkedlist/single"
)

type linkedlist interface {
	Insert(val int)
	Delete(val int)
	Traverse()
	Size() int
}

func main() {
	var singleLink linkedlist
	singleLink = single.NewSingleLinkedList(1)
	singleLink.Insert(2)
	singleLink.Insert(3)
	singleLink.Traverse()
	fmt.Println(singleLink.Size())

	singleLink.Delete(2)
	singleLink.Traverse()
}
