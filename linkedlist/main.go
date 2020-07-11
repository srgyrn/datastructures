package main

import (
	"fmt"
	"github.com/srgyrn/datastructures/linkedlist/single"
)

type Linkedlist interface {
	Insert(val int)
	Delete(val int) error
	Traverse()
	Size() int
}

func main() {
	var singleLink Linkedlist
	singleLink = single.NewLinkedList(1)
	singleLink.Insert(2)
	singleLink.Insert(3)
	singleLink.Traverse()
	fmt.Println(singleLink.Size())

	singleLink.Delete(2)
	singleLink.Traverse()
}
