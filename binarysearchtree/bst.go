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

func (b *bst) search(v int) bool {
	curr := b.root

	for curr != nil {
		if v == curr.val {
			return true
		}

		if v > curr.val {
			curr = curr.right
		} else {
			curr = curr.left
		}
	}

	return false
}

func (b *bst) remove(v int) {
	if !b.search(v) {
		return
	}

	curr := b.root

	for curr != nil {
		if v > curr.val {
			if v == curr.right.val {
				break
			}

			curr = curr.right
		} else {
			if v == curr.left.val {
				break
			}

			curr = curr.left
		}
	}

	isLeaf := func(n *node) bool {
		if n.left == nil && n.right == nil {
			return true
		}
		return false
	}

	hasOneChild := func(n *node) (bool, *node) {
		if n.right != nil && n.left == nil {
			return isLeaf(n.right), n.right
		} else if n.left != nil && n.right == nil {
			return isLeaf(n.left), n.left
		}

		return false, nil
	}

	var vertex *node
	if v > curr.val {
		if isLeaf(curr.right) {
			curr.right = nil
			return
		}
		vertex = curr.right
	} else {
		if isLeaf(curr.left) {
			curr.left = nil
			return
		}
		vertex = curr.left
	}

	// if vertex has one child, assign it to sub-root
	if res, child := hasOneChild(vertex); res {
		*vertex = *child
		return
	}

	//if vertex does not have a right subtree, assign the first greater ancestor its left subtree
	if vertex.right == nil {
		*vertex = *vertex.left
		return
	}

	// if vertex has a right subtree, find min(right subtree) as successor
	curr = vertex.right
	for curr.left.left != nil {
		curr = curr.left
	}

	// set successor val to node
	vertex.val = curr.left.val
	// if successor has a right subtree, assign it to its parent's left
	if curr.left.right != nil {
		curr.left = curr.left.right
	}
}
