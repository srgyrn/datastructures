package binarysearchtree

type node struct {
	val   int
	right *node
	left  *node
}

type bst struct {
	root *node
}

func (b *bst) insert(v int) {
	n := node{val: v}

	// if it's the first item, set it to root
	if b.root == nil {
		b.root = &n
		return
	}

	curr := b.root
	for {
		if curr.val > v {
			if curr.left == nil {
				curr.left = &n
				break
			}

			curr = curr.left
		} else {
			if curr.right == nil {
				curr.right = &n
				break
			}

			curr = curr.right
		}
	}
}