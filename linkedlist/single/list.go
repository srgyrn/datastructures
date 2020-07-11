package single

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func NewLinkedList(val int) *ListNode {
	return &ListNode{
		val:  val,
		next: nil,
	}
}

func (root *ListNode) Insert(val int) {
	newPtr := &ListNode{
		val:  val,
		next: new(ListNode),
	}
	// insert new node at the beginning of the list
	*newPtr.next = *root
	*root = *newPtr
}

func (root *ListNode) Delete(val int) error {
	if root.Size() == 1 {
		return fmt.Errorf("cannot delete root node")
	}

	// if the value is in the root node, just set root to the next node
	if root.val == val {
		*root = *root.next
		return nil
	}

	currentPtr := root
	var prevPtr *ListNode

	for currentPtr.next != nil {
		prevPtr = currentPtr
		currentPtr = currentPtr.next

		// if the value is found, remove it and exit loop
		if currentPtr.val == val {
			prevPtr.next = currentPtr.next
			break
		}
	}

	if currentPtr.val != val {
		return fmt.Errorf("node with value '%d' not found\n", val)
	}

	return nil
}

func (root *ListNode) Traverse() {
	currPtr := root
	for currPtr.next != nil {
		fmt.Printf("%d-->", currPtr.val)
		currPtr = currPtr.next
	}

	fmt.Printf("%d-->nil\n", currPtr.val)
}

func (root *ListNode) Size() int {
	currPtr := root
	count := 1
	for currPtr.next != nil {
		count++
		currPtr = currPtr.next
	}

	return count
}
